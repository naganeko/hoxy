package json

import (
	"hoxy/defs"
	"reflect"
	"strings"
	"sync"
)

// DefMapLock exposes a mutex to allow for safe concurrent access of DefMap.
var DefMapLock = &sync.Mutex{}

// DefMap provides a mapping of an op string to its struct type where possible.
var DefMap = func() map[string]reflect.Type {
	literals := []interface{}{
		defs.CDormGet_build_coin{},
		defs.CDormReceive_favor{},
		defs.CEquipAdjust{},
		defs.CEquipDevelop{},
		defs.CEquipDevelopMulti{},
		defs.CEquipEatEquip{},
		defs.CEquipFinishAllDevelop{},
		defs.CEquipFinishDevelop{},
		defs.CEquipGunEquip{},
		defs.CFairyAdjust{},
		defs.CFairyAdjustConfirm{},
		defs.CFairyEatFairy{},
		defs.CFairyFinishUpgrade{},
		defs.CFairySkillUpgrade{},
		defs.CFairyTeamFairy{},
		defs.CFriendPraise{},
		defs.CFriendUsercard{},
		defs.CFriendVisit{},
		defs.CGunChangeLock{},
		defs.CGunDevelopGun{},
		defs.CGunDevelopLog{},
		defs.CGunDevelopLogCollect{},
		defs.CGunDevelopMultiGun{},
		defs.CGunEatGun{},
		defs.CGunFinishAllDevelop{},
		defs.CGunFinishDevelop{},
		defs.CGunFinishUpgrade{},
		defs.CGunFixFinish{},
		defs.CGunFixGuns{},
		defs.CGunQuickDevelop{},
		defs.CGunRetireGun{},
		defs.CGunSetPosition{},
		defs.CGunSkillUpgrade{},
		defs.CGunTeamGun{},
		defs.CIndexGetMailList{},
		defs.CIndexGetResourceInMail{},
		defs.CIndexGuide{},
		defs.CIndexHome{},
		defs.CIndexIndex{},
		defs.CIndexQuickGetQuestsResourceInMails{},
		defs.CMissionBattleFinish{},
		defs.CMissionCoinBattleFinish{},
		defs.CMissionEndTrialExercise{},
		defs.CMissionFriendTeamAi{},
		defs.CMissionReinforceFriendTeam{},
		defs.CMissionReinforceTeam{},
		defs.CMissionStartMission{},
		defs.CMissionStartTrial{},
		defs.CMissionSupplyTeam{},
		defs.CMissionTeamMove{},
		defs.COperationFinishOperation{},
		defs.COperationStartOperation{},
		defs.COuthouseEstablish_build_finish{},
		defs.COuthouseEstablish_build{},
		defs.SDormGet_build_coin{},
		defs.SDormKalinaFavor(0),
		defs.SDormReceive_favor{},
		defs.SDormShare{},
		defs.SEquipAdjust{},
		defs.SEquipDevelop{},
		defs.SEquipDevelopCollectList{},
		defs.SEquipDevelopLog{},
		defs.SEquipDevelopMulti{},
		defs.SEquipEatEquip{},
		defs.SEquipFinishAllDevelop{},
		defs.SEquipFinishDevelop{},
		defs.SEquipGunEquip(0),
		defs.SFairyAdjust{},
		defs.SFairyAdjustConfirm(0),
		defs.SFairyEatFairy{},
		defs.SFairyFinishUpgrade(0),
		defs.SFairySkillUpgrade(0),
		defs.SFairyTeamFairy(0),
		defs.SFriendApplylist{},
		defs.SFriendDormInfo{},
		defs.SFriendList{},
		defs.SFriendMessagelist{},
		defs.SFriendPraise{},
		defs.SFriendRandomList{},
		defs.SFriendTeamGuns{},
		defs.SFriendUsercard{},
		defs.SFriendVisit{},
		defs.SGunChangeLock(0),
		defs.SGunDevelopCollectList{},
		defs.SGunDevelopGun{},
		defs.SGunDevelopLog{},
		defs.SGunDevelopLogCollect(0),
		defs.SGunDevelopMultiGun{},
		defs.SGunEatGun{},
		defs.SGunFinishAllDevelop{},
		defs.SGunFinishDevelop{},
		defs.SGunFinishUpgrade(0),
		defs.SGunFixFinish(0),
		defs.SGunFixGuns(0),
		defs.SGunRetireGun(0),
		defs.SGunSetPosition(0),
		defs.SGunSkillUpgrade(0),
		defs.SGunTeamGun(0),
		defs.SIndexAttendance{},
		defs.SIndexDownloadSuccess(0),
		defs.SIndexGetMailList{},
		defs.SIndexGetResourceInMail{},
		defs.SIndexGetUidEnMicaQueue{},
		defs.SIndexGuide(0),
		defs.SIndexHome{},
		defs.SIndexIndex{},
		defs.SIndexQuest{},
		defs.SIndexQuickGetQuestsResourceInMails{},
		defs.SIndexQuickGetResourceInMails{},
		defs.SIndexRecoverBp{},
		defs.SIndexVersion{},
		defs.SMissionBattleFinish{},
		defs.SMissionCoinBattleFinish{},
		defs.SMissionDrawEvent{},
		defs.SMissionEndEnemyTurn{},
		defs.SMissionEndTrialExercise{},
		defs.SMissionEndTurn{},
		defs.SMissionFriendTeamAi(0),
		defs.SMissionFriendTeamMove{},
		defs.SMissionReinforceFriendTeam{},
		defs.SMissionReinforceTeam{},
		defs.SMissionStartMission{},
		defs.SMissionStartTrial{},
		defs.SMissionStartTurn{},
		defs.SMissionSupplyTeam(0),
		defs.SMissionTeamMove{},
		defs.SOperationFinishOperation{},
		defs.SOperationStartOperation(0),
		defs.SOuthouseEstablish_build_finish{},
		defs.SOuthouseEstablish_build{},
	}
	ret := make(map[string]reflect.Type)
	for _, def := range literals {
		t := reflect.TypeOf(def)
		op := defTypeToOpString(t)
		op = strings.ToLower(op)
		ret[op] = t
	}
	return ret
}()
