package proxy

import (
	"github.com/kyoukaya/hoxy/defs"
	"github.com/kyoukaya/hoxy/log"
	"github.com/kyoukaya/hoxy/proxy/json"
	"net/http"

	"github.com/elazarl/goproxy"
)

func (hoxy *HoxyProxy) dispatchReq(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	// reqID := req.FormValue("req_id")
	reqURL := req.URL.String()
	op := "C" + reqURL[46:]
	userCtx := hoxy.GetUser(req.FormValue("uid"))
	reqCtx := GetDispatchContext(ctx)
	reqCtx.RequestOp = op

	var dec []byte
	var err error
	var isSignCode bool
	// A request will either contain an outdatacode field or a signcode field but not both.
	if enc := req.FormValue("outdatacode"); enc != "" {
		dec, _, err = userCtx.Decode(enc)
	} else if enc := req.FormValue("signcode"); enc != "" {
		dec, _, err = userCtx.Decode(enc)
		isSignCode = true
	} else {
		// These packets don't have a signcode or outdatacode field.
		if op != "CIndex/getUidEnMicaQueue" && op != "CIndex/version" {
			log.Warnf(">>## Something weird happened while decoding %s\n %#v\n", op, req.Form)
		}
		dec = nil
	}

	if err != nil {
		log.Warnf(">>## [%s] failed to decode for UID: %s\n", op, req.Form.Get("uid"))
		return req, nil
	}

	if len(dec) == 0 {
		log.Infof(">>## [%s]", op)
	} else {
		log.Infof(">>## [%s]", op)
	}

	if isSignCode {
		dec = nil
	}
	return hoxy.dispatch(op, dec, ctx)
}

func (hoxy *HoxyProxy) dispatch(op string, dec []byte, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	var unmarshalErr error
	var req = ctx.Req
	var res = ctx.Resp

	user := hoxy.GetUser(req.FormValue("uid"))
	// Run raw hooks
	if user != nil {
		if hooks, ok := user.RawHooks["*"]; ok {
			for _, hook := range hooks {
				hook.Handle(op, dec, user, ctx)
			}
		}
		if hooks, ok := user.RawHooks[op]; ok {
			for _, hook := range hooks {
				hook.Handle(op, dec, user, ctx)
			}
		}
	}

	pkt, _, unmarshalErr := json.UnMarshal(op, dec)
	if unmarshalErr != nil {
		switch unmarshalErr.(type) {
		case json.MarshalNoDefErr:
			e := unmarshalErr.(json.MarshalNoDefErr)
			log.Warnf("No definitions found for %s", e.Op)
		case json.MarshalMismatchErr:
			e := unmarshalErr.(json.MarshalMismatchErr)
			log.Warnf("Marshal->Unmarshal mismatch for %s", e.Op)
		default:
			log.Warnln(unmarshalErr)
			return req, res
		}
	}
	// Include the unmarshalled packet into the request context if successfully unmarshalled
	dispatchCtx := GetDispatchContext(ctx)
	if ctx.Resp == nil {
		dispatchCtx.RequestUnmarshalErr = unmarshalErr
		dispatchCtx.RequestPkt = pkt
	} else {
		dispatchCtx.ResponseUnmarshalErr = unmarshalErr
	}

	// Run auth hook if SIndex/getUidEnMicaQueue is unmarshalled.
	if _, ok := pkt.(*defs.SIndexGetUidEnMicaQueue); ok && unmarshalErr == nil {
		openID, UID, sign, longtoken, err := hoxy.authHandler(op, pkt, ctx)
		if err != nil {
			log.Warnf("AuthHook failed to initialize ctx: %s", err)
			return req, res
		}
		hoxy.addUser(openID, UID, sign, longtoken)
		// UserCtx should be initialized by the AuthHook hook
		user := hoxy.GetUser(pkt.(*defs.SIndexGetUidEnMicaQueue).UID)
		user.initMods(modules)
	}

	// Run packet hooks if user is initialized
	if user := hoxy.GetUser(ctx.Req.FormValue("uid")); user != nil {
		// Run wildcard hooks
		if hooks, ok := user.Hooks["*"]; ok {
			for _, hook := range hooks {
				_, err := hook.Handle(op, pkt, user, ctx)
				if err != nil {
					log.Warnf("dispatch: error occured while running hook %s on packet %s: %s\n"+
						"data: %s", hook.name, op, err.Error(), dec)
				}
			}
		}
		// Run normal hooks for op
		if hooks, ok := user.Hooks[op]; ok {
			for _, hook := range hooks {
				_, err := hook.Handle(op, pkt, user, ctx)
				if err != nil {
					log.Warnf("dispatch: error occured while running hook %s on packet %s: %s\n"+
						"data: %s", hook.name, op, err.Error(), dec)
				}
			}
		}
	}

	return req, res
}

func (hoxy *HoxyProxy) dispatchRes(body []byte, resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	var err error
	var dec []byte
	if len(body) == 0 {
		return resp
	}

	op := "S" + ctx.Req.URL.String()[46:]
	if body[0] == '#' {
		user := hoxy.GetUser(ctx.Req.FormValue("uid"))
		dec, _, err = user.Decode(string(body))
		if err != nil {
			log.Warnf("<<## [%s] failed to decode for UID: %s\n", op, ctx.Req.Form.Get("uid"))
			return resp
		}
		log.Infof("<<## [%s]", op)
	} else {
		log.Infof("<<<< [%s]", op)
		// dispatch packets like SIndex/index which are not encrypted
		dec = body
	}
	_, resp = hoxy.dispatch(op, dec, ctx)

	return resp
}
