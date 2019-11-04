package settinginfo

// ConfigInfo setting struct
type ConfigInfo struct {
	ULGLoginURL             string `json:"ULGLoginURL"`
	ULGGetuserURL           string `json:"ULGGetuserURL"`
	ULGAuthorizedURL        string `json:"ULGAuthorizedURL"`
	ULGExchangeURL          string `json:"ULGExchangeURL"`
	ULGCheckoutURL          string `json:"ULGCheckoutURL"`
	GameTypeID              string `json:"GameTypeID"`
	IP                      string `json:"IP"`
	PORT                    string `json:"PORT"`
	DBIP                    string `json:"DBIP"`
	DBPORT                  string `json:"DBPORT"`
	DBUser                  string `json:"DBUser"`
	DBPassword              string `json:"DBPassword"`
	AccountEncodeStr        string `json:"AccountEncodeStr"`
	RedisURL                string `json:"RedisURL"`
	UpdateDBTime            string `json:"UpdateDBTime"`
	MaintainStartTime       string `json:"MaintainStartTime"`
	MaintainFinishTime      string `json:"MaintainFinishTime"`
	ULGMaintainCheckoutTime string `json:"ULGMaintainCheckoutTime"`
	RTPSetting              int    `json:"RTPSetting"`
	RespinSetting           int    `json:"RespinSetting"`
	ServerDayPayLimit       int64  `json:"ServerDayPayLimit"`
	ServerDayPayDefault     int64  `json:"ServerDayPayDefault"`
	WinScoreLimit           int64  `json:"WinScoreLimit"`
	WinBetRateLimit         int64  `json:"WinBetRateLimit"`
	Maintain                bool   `json:"Maintain"`
	DebugLog                bool   `json:"DebugLog"`
}

// // GetGameTypeID ...
// func (C *ConfigInfo) GetGameTypeID() string {
// 	return C.GameTypeID
// }

// // GetRespinSetting ...
// func (C *ConfigInfo) GetRespinSetting() int {
// 	return C.RespinSetting
// }

// // GetRTPSetting ...
// func (C *ConfigInfo) GetRTPSetting() int {
// 	return C.RTPSetting
// }

// // GetWinScoreLimit ...
// func (C *ConfigInfo) GetWinScoreLimit() int64 {
// 	return C.WinScoreLimit
// }

// // GetWinBetRateLimit ...
// func (C *ConfigInfo) GetWinBetRateLimit() int64 {
// 	return C.WinBetRateLimit
// }

// // GetMaintainCheckoutTime ...
// func (C *ConfigInfo) GetMaintainCheckoutTime() string {
// 	return C.ULGMaintainCheckoutTime
// }

// // GetLoginURL ...
// func (C *ConfigInfo) GetLoginURL() string {
// 	return C.ULGLoginURL
// }

// // GetUserURL ...
// func (C *ConfigInfo) GetUserURL() string {
// 	return C.ULGLoginURL
// }

// // GetAuthorizedURL ...
// func (C *ConfigInfo) GetAuthorizedURL() string {
// 	return C.ULGAuthorizedURL
// }

// // GetCheckoutURL ...
// func (C *ConfigInfo) GetCheckoutURL() string {
// 	return C.ULGCheckoutURL
// }

// // GameRuleInfo ...
// type GameRuleInfo interface {
// 	GetGameTypeID() string
// 	GetRespinSetting() int
// 	GetRTPSetting() int
// 	GetWinScoreLimit() int64
// 	GetWinBetRateLimit() int64
// }

// // UlgAPIInfo ...
// type UlgAPIInfo interface {
// 	GetMaintainCheckoutTime() string
// 	GetLoginURL() string
// 	GetUserURL() string
// 	GetAuthorizedURL() string
// 	GetExchangeURL() string
// 	GetCheckoutURL() string
// }
