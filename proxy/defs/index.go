package defs

// CIndexIndex Client's request for Index/index
type CIndexIndex struct {
	Time int `json:"time"`
}

// SIndexIndex Server's response for Index/index
type SIndexIndex struct {
	UserInfo struct {
		Gem                         int    `json:"gem"`
		ID                          int    `json:"id"`
		UserID                      int    `json:"user_id"`
		OpenID                      int    `json:"open_id"`
		ChannelID                   int    `json:"channel_id"`
		Sign                        string `json:"sign"`
		Name                        string `json:"name"`
		GemPay                      int    `json:"gem_pay"`
		PauseTurnChance             int    `json:"pause_turn_chance"`
		Experience                  int    `json:"experience"`
		Lv                          int    `json:"lv"`
		Maxequip                    int    `json:"maxequip"`
		Maxgun                      int    `json:"maxgun"`
		Maxteam                     int    `json:"maxteam"`
		MaxBuildSlot                int    `json:"max_build_slot"`
		MaxEquipBuildSlot           int    `json:"max_equip_build_slot"`
		MaxFixSlot                  int    `json:"max_fix_slot"`
		MaxUpgradeSlot              int    `json:"max_upgrade_slot"`
		MaxGunPreset                int    `json:"max_gun_preset"`
		MaxFairy                    int    `json:"max_fairy"`
		Bp                          int    `json:"bp"`
		BpPay                       int    `json:"bp_pay"`
		Mp                          int    `json:"mp"`
		Ammo                        int    `json:"ammo"`
		Mre                         int    `json:"mre"`
		Part                        int    `json:"part"`
		Core                        int    `json:"core"`
		Coin1                       int    `json:"coin1"`
		Coin2                       int    `json:"coin2"`
		Coin3                       int    `json:"coin3"`
		Monthlycard1EndTime         int    `json:"monthlycard1_end_time"`
		Monthlycard2EndTime         int    `json:"monthlycard2_end_time"`
		Monthlycard3EndTime         int    `json:"monthlycard3_end_time"`
		Growthfond                  int    `json:"growthfond"`
		LastRecoverTime             int    `json:"last_recover_time"`
		LastBpRecoverTime           int    `json:"last_bp_recover_time"`
		LastFavorRecoverTime        int    `json:"last_favor_recover_time"`
		LastMonthlycard1RecoverTime int    `json:"last_monthlycard1_recover_time"`
		LastMonthlycard2RecoverTime int    `json:"last_monthlycard2_recover_time"`
		LastLoginTime               int    `json:"last_login_time"`
		RegTime                     int    `json:"reg_time"`
		GunCollect                  string `json:"gun_collect"`
		Maxdorm                     int    `json:"maxdorm"`
		MaxdormSpare                int    `json:"maxdorm_spare"`
		LastDormMaterialChangeTime1 int    `json:"last_dorm_material_change_time1"`
		LastDormMaterialChangeTime2 int    `json:"last_dorm_material_change_time2"`
		MaterialAvailableNum1       int    `json:"material_available_num1"`
		MaterialAvailableNum2       int    `json:"material_available_num2"`
		DormMaxScore                int    `json:"dorm_max_score"`
		LastSsocChangeTime          int    `json:"last_ssoc_change_time"`
		AppGuardID                  string `json:"app_guard_id"`
		AppGuardNum                 int    `json:"app_guard_num"`
		IsBind                      bool   `json:"is_bind"`
	} `json:"user_info"`
	UpgradeActInfo []struct {
		UserID          string `json:"user_id"`
		FairyWithUserID string `json:"fairy_with_user_id"`
		GunWithUserID   string `json:"gun_with_user_id"`
		Type            string `json:"type"`
		Skill           string `json:"skill"`
		UpgradeSlot     string `json:"upgrade_slot"`
		EndTime         string `json:"end_time"`
	} `json:"upgrade_act_info"`
	GuideInfo struct {
		ID     int    `json:"id"`
		UserID int    `json:"user_id"`
		Guide  string `json:"guide"`
	} `json:"guide_info"`
	ItemWithUserInfo []struct {
		ItemID int `json:"item_id"`
		Number int `json:"number"`
		ID     int `json:"id"`
	} `json:"item_with_user_info"`
	ItemLimitWithUserType9Info []struct {
		ItemID         int `json:"item_id"`
		DailyGet       int `json:"daily_get"`
		DailyClearTime int `json:"daily_clear_time"`
	} `json:"item_limit_with_user_type9_info"`
	MissionWithUserInfo []struct {
		ID                    int    `json:"id"`
		UserID                int    `json:"user_id"`
		MissionID             int    `json:"mission_id"`
		Medal1                int    `json:"medal1"`
		Medal2                int    `json:"medal2"`
		Bestrank              int    `json:"bestrank"`
		Medal4                int    `json:"medal4"`
		Counter               int    `json:"counter"`
		WinCounter            int    `json:"win_counter"`
		ShortestInCoinmission string `json:"shortest_in_coinmission"`
		Type5Score            int    `json:"type5_score"`
		IsOpen                int    `json:"is_open"`
		IsDropDrawEvent       int    `json:"is_drop_draw_event"`
		IsClose               int    `json:"is_close"`
		CycleWinCount         int    `json:"cycle_win_count"`
		MappedWinCounter      int    `json:"mapped_win_counter"`
	} `json:"mission_with_user_info"`
	OperationActInfo []struct {
		ID          int `json:"id"`
		OperationID int `json:"operation_id"`
		UserID      int `json:"user_id"`
		TeamID      int `json:"team_id"`
		StartTime   int `json:"start_time"`
	} `json:"operation_act_info"`
	DevelopActInfo []struct {
		ID         int `json:"id"`
		BuildSlot  int `json:"build_slot"`
		UserID     int `json:"user_id"`
		GunID      int `json:"gun_id"`
		EquipID    int `json:"equip_id"`
		StartTime  int `json:"start_time"`
		Mp         int `json:"mp"`
		Ammo       int `json:"ammo"`
		Mre        int `json:"mre"`
		Part       int `json:"part"`
		InputLevel int `json:"input_level"`
		Core       int `json:"core"`
		Item1Num   int `json:"item1_num"`
	} `json:"develop_act_info"`
	FixActInfo     []interface{} `json:"fix_act_info"`
	MissionActInfo interface{}   `json:"mission_act_info"`
	SpotActInfo    interface{}   `json:"spot_act_info"`
	UserRecord     struct {
		UserID                         int     `json:"user_id"`
		MissionCampaign                int     `json:"mission_campaign"`
		SpecialMissionCampaign         int     `json:"special_mission_campaign"`
		AttendanceType1Day             int     `json:"attendance_type1_day"`
		AttendanceType1Time            int     `json:"attendance_type1_time"`
		AttendanceType2Day             int     `json:"attendance_type2_day"`
		AttendanceType2Time            int     `json:"attendance_type2_time"`
		DevelopLowrankCount            int     `json:"develop_lowrank_count"`
		SpecialDevelopLowrankCount     int     `json:"special_develop_lowrank_count"`
		GetGiftIds                     string  `json:"get_gift_ids"`
		SpendPoint                     string  `json:"spend_point"`
		GemMallIds                     string  `json:"gem_mall_ids"`
		ProductType21Ids               string  `json:"product_type21_ids"`
		SevenType                      int     `json:"seven_type"`
		SevenStartTime                 int     `json:"seven_start_time"`
		SevenAttendanceDays            string  `json:"seven_attendance_days"`
		SevenSpendPoint                int     `json:"seven_spend_point"`
		LastBpBuyTime                  int     `json:"last_bp_buy_time"`
		BpBuyCount                     int     `json:"bp_buy_count"`
		NewDevelopgunCount             int     `json:"new_developgun_count"`
		Buyitem1DevelopgunCount        float64 `json:"buyitem1_developgun_count"`
		Buyitem1SpecialdevelopgunCount string  `json:"buyitem1_specialdevelopgun_count"`
		Buyitem1Num                    int     `json:"buyitem1_num"`
		LastDevelopgunTime             int     `json:"last_developgun_time"`
		LastSpecialdevelopgunTime      int     `json:"last_specialdevelopgun_time"`
		FurnitureClasses               string  `json:"furniture_classes"`
		Adjutant                       string  `json:"adjutant"`
		AdjutantFairy                  string  `json:"adjutant_fairy"`
		MissionGroupTodayResetNum      int     `json:"mission_group_today_reset_num"`
		MissionGroupLastResetTime      int     `json:"mission_group_last_reset_time"`
		SpendpointAgeID                int     `json:"spendpoint_age_id"`
		SpendPointThismonth            int     `json:"spend_point_thismonth"`
		LastSpendpointTime             int     `json:"last_spendpoint_time"`
		NextCoreRecoverGunTime         int     `json:"next_core_recover_gun_time"`
		LastDatacellChangeTime         int     `json:"last_datacell_change_time"`
		BuympNum                       int     `json:"buymp_num"`
		BuyammoNum                     int     `json:"buyammo_num"`
		BuymreNum                      int     `json:"buymre_num"`
		BuypartNum                     int     `json:"buypart_num"`
		DevelopNonewNum                int     `json:"develop_nonew_num"`
	} `json:"user_record"`
	EquipWithUserInfo map[int]struct {
		ID               int    `json:"id"`
		UserID           int    `json:"user_id"`
		GunWithUserID    int    `json:"gun_with_user_id"`
		EquipID          int    `json:"equip_id"`
		EquipExp         int    `json:"equip_exp"`
		EquipLevel       int    `json:"equip_level"`
		Pow              int    `json:"pow"`
		Hit              int    `json:"hit"`
		Dodge            int    `json:"dodge"`
		Speed            int    `json:"speed"`
		Rate             int    `json:"rate"`
		CriticalHarmRate int    `json:"critical_harm_rate"`
		CriticalPercent  int    `json:"critical_percent"`
		ArmorPiercing    int    `json:"armor_piercing"`
		Armor            int    `json:"armor"`
		Shield           int    `json:"shield"`
		DamageAmplify    int    `json:"damage_amplify"`
		DamageReduction  int    `json:"damage_reduction"`
		NightViewPercent int    `json:"night_view_percent"`
		BulletNumberUp   int    `json:"bullet_number_up"`
		AdjustCount      int    `json:"adjust_count"`
		IsLocked         int    `json:"is_locked"`
		LastAdjust       string `json:"last_adjust"`
	} `json:"equip_with_user_info"`
	DevelopEquipActInfo map[int]struct {
		Type       int `json:"type"`
		EquipID    int `json:"equip_id"`
		BuildSlot  int `json:"build_slot"`
		UserID     int `json:"user_id"`
		StartTime  int `json:"start_time"`
		Mp         int `json:"mp"`
		Ammo       int `json:"ammo"`
		Mre        int `json:"mre"`
		Part       int `json:"part"`
		InputLevel int `json:"input_level"`
		Core       int `json:"core"`
		ItemNum    int `json:"item_num"`
	} `json:"develop_equip_act_info"`
	TeamInCoinMissionInfo map[int]struct {
		UserID        string `json:"user_id"`
		Location      string `json:"location"`
		GunWithUserID string `json:"gun_with_user_id"`
		Position      string `json:"position"`
	} `json:"team_in_coin_mission_info"`
	SkinWithUserInfo map[int]struct {
		SkinID int  `json:"skin_id"`
		UserID int  `json:"user_id"`
		IsRead bool `json:"is_read"`
	} `json:"skin_with_user_info"`
	EventInfo struct {
		Num315 struct {
			ID        string `json:"id"`
			Code      string `json:"code"`
			Condition string `json:"condition"`
			StartTime int    `json:"start_time"`
			EndTime   int    `json:"end_time"`
		} `json:"315"`
	} `json:"event_info"`
	DormScoreInfo map[int]struct {
		UserID string `json:"user_id"`
		DormID string `json:"dorm_id"`
		Score  string `json:"score"`
	} `json:"dorm_score_info"`
	FurnitureWithUserInfo []struct {
		ID          int `json:"id"`
		Dorm        int `json:"dorm"`
		FurnitureID int `json:"furniture_id"`
		X           int `json:"x"`
		Y           int `json:"y"`
	} `json:"furniture_with_user_info"`
	FurnitureCollectInfo []struct {
		FurnitureID int `json:"furniture_id"`
		UserID      int `json:"user_id"`
	} `json:"furniture_collect_info"`
	FairyCollectInfo []struct {
		FairyID int `json:"fairy_id"`
		UserID  int `json:"user_id"`
	} `json:"fairy_collect_info"`
	TrialWithUserInfo struct {
		Counter  int `json:"counter"`
		UserID   int `json:"user_id"`
		MaxTrial int `json:"max_trial"`
		MaxTime  int `json:"max_time"`
	} `json:"trial_with_user_info"`
	TrialExerciseInfo struct {
		GunsInfo             string `json:"guns_info"`
		BattleTeam           int    `json:"battle_team"`
		LastBattleFinishTime int    `json:"last_battle_finish_time"`
		Cheat                int    `json:"cheat"`
		TrialID              int    `json:"trial_id"`
		TeamIds              int    `json:"team_ids"`
		UserID               int    `json:"user_id"`
	} `json:"trial_exercise_info"`
	GunPresetInfo []struct {
		UserID   int    `json:"user_id"`
		PresetNo int    `json:"preset_no"`
		Gun1     string `json:"gun1"`
		Gun2     string `json:"gun2"`
		Gun3     string `json:"gun3"`
		Gun4     string `json:"gun4"`
		Gun5     string `json:"gun5"`
	} `json:"gun_preset_info"`
	OuthouseEstablishInfo []struct {
		ID                 int         `json:"id"`
		UserID             int         `json:"user_id"`
		RoomID             int         `json:"room_id"`
		EstablishID        int         `json:"establish_id"`
		EstablishType      int         `json:"establish_type"`
		FurnitureID        int         `json:"furniture_id"`
		UpgradeEstablishID int         `json:"upgrade_establish_id"`
		UpgradeStarttime   int         `json:"upgrade_starttime"`
		BuildStarttime     int         `json:"build_starttime"`
		BuildNum           int         `json:"build_num"`
		BuildTmpData       interface{} `json:"build_tmp_data"`
		BuildGetTime       int         `json:"build_get_time"`
		UpdateFurnitureID  int         `json:"update_furniture_id"`
		FurniturePostion   string      `json:"furniture_postion"`
		EstablishLv        int         `json:"establish_lv"`
		UpgradeCoin        int         `json:"upgrade_coin"`
		UpgradeTime        int         `json:"upgrade_time"`
		UpgradeCondition   string      `json:"upgrade_condition"`
		Parameter1         string      `json:"parameter_1"`
		Parameter2         string      `json:"parameter_2"`
		Parameter3         string      `json:"parameter_3"`
	} `json:"outhouse_establish_info"`
	DormRestFriendBuildCoinCount int `json:"dorm_rest_friend_build_coin_count"`
	FairyWithUserInfo            map[int]struct {
		ID                  int    `json:"id"`
		UserID              int    `json:"user_id"`
		FairyID             int    `json:"fairy_id"`
		TeamID              int    `json:"team_id"`
		FairyLv             int    `json:"fairy_lv"`
		FairyExp            int    `json:"fairy_exp"`
		QualityLv           int    `json:"quality_lv"`
		QualityExp          int    `json:"quality_exp"`
		SkillLv             int    `json:"skill_lv"`
		PassiveSkill        int    `json:"passive_skill"`
		IsLocked            int    `json:"is_locked"`
		EquipID             int    `json:"equip_id"`
		AdjustCount         int    `json:"adjust_count"`
		LastAdjust          int    `json:"last_adjust"`
		PassiveSkillCollect string `json:"passive_skill_collect"`
		Skin                string `json:"skin"`
	} `json:"fairy_with_user_info"`
	TipsInfo              interface{} `json:"tips_info"`
	FairySkinWithUserInfo map[int]struct {
		ID      int `json:"id"`
		SkinID  int `json:"skin_id"`
		UserID  int `json:"user_id"`
		FairyID int `json:"fairy_id"`
	} `json:"fairy_skin_with_user_info"`
	UserGameSetting struct {
		UID      int    `json:"uid"`
		Settings string `json:"settings"`
	} `json:"user_game_setting"`
	MissionKeyWithUserInfo map[int]int `json:"mission_key_with_user_info"`
	SquadDataDaily         []struct {
		UserID         int    `json:"user_id"`
		SquadID        int    `json:"squad_id"`
		Type           string `json:"type"`
		LastFinishTime string `json:"last_finish_time"`
		Count          int    `json:"count"`
		Receive        int    `json:"receive"`
	} `json:"squad_data_daily"`
	SquadWithUserInfo   []interface{} `json:"squad_with_user_info"`
	ChipWithUserInfo    []interface{} `json:"chip_with_user_info"`
	DataAnalysisActInfo []interface{} `json:"data_analysis_act_info"`
	SquadTrainActInfo   []interface{} `json:"squad_train_act_info"`
	SquadSkillActInfo   []interface{} `json:"squad_skill_act_info"`
	SquadFixActInfo     []interface{} `json:"squad_fix_act_info"`
	SurplusDatacellInfo struct {
		UserID int `json:"user_id"`
		Base   int `json:"base"`
		Senior int `json:"senior"`
	} `json:"surplus_datacell_info"`
	GunMemoirList       []interface{} `json:"gun_memoir_list"`
	GunInTheaterInfo    []interface{} `json:"gun_in_theater_info"`
	SquadInTheaterInfo  []interface{} `json:"squad_in_theater_info"`
	FairyInTheaterInfo  []interface{} `json:"fairy_in_theater_info"`
	TheaterExerciseInfo interface{}   `json:"theater_exercise_info"`
	TheaterWithUserInfo []interface{} `json:"theater_with_user_info"`
	FriendWithUserInfo  []struct {
		Type          int    `json:"type"`
		EndTime       int    `json:"end_time"`
		FUserid       int    `json:"f_userid"`
		Name          string `json:"name"`
		Lv            int    `json:"lv"`
		HeadpicID     int    `json:"headpic_id"`
		HomepageTime  int    `json:"homepage_time"`
		GiveTeamToday int    `json:"give_team_today"`
	} `json:"friend_with_user_info"`
	UserFriendInfo struct {
		FUserid int    `json:"f_userid"`
		Name    string `json:"name"`
		Lv      int    `json:"lv"`
		Medal   []struct {
			ID    int `json:"id"`
			Level int `json:"level"`
			Num   int `json:"num"`
		} `json:"medal"`
		CardID    int    `json:"card_id"`
		HeadpicID int    `json:"headpic_id"`
		HomebgID  int    `json:"homebg_id"`
		Intro     string `json:"intro"`
		TeamInfo  []struct {
			ID             int  `json:"id"`
			UserID         int  `json:"user_id"`
			GroupID        int  `json:"group_id"`
			GunID          int  `json:"gun_id"`
			GunExp         int  `json:"gun_exp"`
			GunLevel       int  `json:"gun_level"`
			Location       int  `json:"location"`
			Position       int  `json:"position"`
			Pow            int  `json:"pow"`
			Hit            int  `json:"hit"`
			Dodge          int  `json:"dodge"`
			Rate           int  `json:"rate"`
			Skill1         int  `json:"skill1"`
			Skill2         int  `json:"skill2"`
			Number         int  `json:"number"`
			Equip1         int  `json:"equip1"`
			Equip2         int  `json:"equip2"`
			Equip3         int  `json:"equip3"`
			Favor          int  `json:"favor"`
			Skin           int  `json:"skin"`
			SoulBond       bool `json:"soul_bond"`
			IfModification int  `json:"if_modification"`
			GunWithUserID  int  `json:"gun_with_user_id"`
		} `json:"team_info"`
		BorrowTeamToday int `json:"borrow_team_today"`
		FairyInfo       []struct {
			ID              int `json:"id"`
			UserID          int `json:"user_id"`
			GroupID         int `json:"group_id"`
			FairyID         int `json:"fairy_id"`
			FairyLv         int `json:"fairy_lv"`
			FairyExp        int `json:"fairy_exp"`
			QualityLv       int `json:"quality_lv"`
			QualityExp      int `json:"quality_exp"`
			SkillLv         int `json:"skill_lv"`
			PassiveSkill    int `json:"passive_skill"`
			EquipID         int `json:"equip_id"`
			Skin            int `json:"skin"`
			FairyWithUserID int `json:"fairy_with_user_id"`
		} `json:"fairy_info"`
		IsReturnUser int `json:"is_return_user"`
	} `json:"user_friend_info"`
	GiftWithUserInfo []struct {
		ItemID int `json:"item_id"`
		Number int `json:"number"`
	} `json:"gift_with_user_info"`
	MedalWithUserInfo []struct {
		MedalID int `json:"medal_id"`
		Num     int `json:"num"`
	} `json:"medal_with_user_info"`
	KalinaWithUserInfo struct {
		UserID       int `json:"user_id"`
		ClickNum     int `json:"click_num"`
		ClickTime    int `json:"click_time"`
		Level        int `json:"level"`
		Favor        int `json:"favor"`
		LastFavor    int `json:"last_favor"`
		Skin         int `json:"skin"`
		SendMailTime int `json:"send_mail_time"`
	} `json:"kalina_with_user_info"`
	ShareWithUserInfo struct {
		LastTime int `json:"last_time"`
	} `json:"share_with_user_info"`
	AutoMissionActInfo interface{} `json:"auto_mission_act_info"`
	GunWithUserInfo    []struct {
		ID             int  `json:"id"`
		UserID         int  `json:"user_id"`
		GunID          int  `json:"gun_id"`
		GunExp         int  `json:"gun_exp"`
		GunLevel       int  `json:"gun_level"`
		TeamID         int  `json:"team_id"`
		IfModification bool `json:"if_modification"`
		Location       int  `json:"location"`
		Position       int  `json:"position"`
		Life           int  `json:"life"`
		Ammo           int  `json:"ammo"`
		Mre            int  `json:"mre"`
		Pow            int  `json:"pow"`
		Hit            int  `json:"hit"`
		Dodge          int  `json:"dodge"`
		Rate           int  `json:"rate"`
		Skill1         int  `json:"skill1"`
		Skill2         int  `json:"skill2"`
		FixEndTime     int  `json:"fix_end_time"`
		IsLocked       int  `json:"is_locked"`
		Number         int  `json:"number"`
		Equip1         int  `json:"equip1"`
		Equip2         int  `json:"equip2"`
		Equip3         int  `json:"equip3"`
		Equip4         int  `json:"equip4"`
		Favor          int  `json:"favor"`
		MaxFavor       int  `json:"max_favor"`
		FavorToplimit  int  `json:"favor_toplimit"`
		SoulBond       bool `json:"soul_bond"`
		Skin           int  `json:"skin"`
		CanClick       bool `json:"can_click"`
		SoulBondTime   int  `json:"soul_bond_time"`
	} `json:"gun_with_user_info"`
	ShowCdkey        int `json:"show_cdkey"`
	CanUsePayceo     int `json:"can_use_payceo"`
	Hexie            int `json:"hexie"`
	Nf               int `json:"nf"`
	IosAds           int `json:"ios_ads"`
	MissionEventInfo struct {
		ID                   int    `json:"id"`
		MissionCampaign      int    `json:"mission_campaign"`
		DrawEventID          string `json:"draw_event_id"`
		StartTime            int    `json:"start_time"`
		EndTime              int    `json:"end_time"`
		NormalCombatCampaign string `json:"normal_combat_campaign"`
		NormalCombatInit     string `json:"normal_combat_init"`
		DrawInfo             []struct {
			DrawEventID int `json:"draw_event_id"`
			DrawNum     int `json:"draw_num"`
			DrawEndtime int `json:"draw_endtime"`
		} `json:"draw_info"`
		ItemLimitWithUser map[int]struct {
			ItemID         string `json:"item_id"`
			UserID         string `json:"user_id"`
			DailyGet       string `json:"daily_get"`
			DailyClearTime string `json:"daily_clear_time"`
			EventGet       string `json:"event_get"`
			EndTime        string `json:"end_time"`
		} `json:"item_limit_with_user"`
		TomorrowUnix int `json:"tomorrow_unix"`
		AfterDays    int `json:"after_days"`
	} `json:"mission_event_info"`
	AuthenticationURL    string `json:"authentication_url"`
	NaiveBuildGunFormula string `json:"naive_build_gun_formula"`
	GameConfigInfo       struct {
		ShareGlobalSwitch    bool `json:"share_global_switch"`
		SpecialDevelopSwitch bool `json:"special_develop_switch"`
		EquipEnhanceSwitch   bool `json:"equip_enhance_switch"`
		EquipRectifySwitch   bool `json:"equip_rectify_switch"`
		TrailSwitch          bool `json:"trail_switch"`
		SoulbindGlobalSwitch bool `json:"soulbind_global_switch"`
	} `json:"game_config_info"`
}
