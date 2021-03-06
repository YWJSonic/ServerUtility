package ulginfo

import (
	"strconv"

	"github.com/YWJSonic/ServerUtility/foundation"
	"github.com/YWJSonic/ServerUtility/playerinfo"
)

// Info plant ULG game play data
type Info struct {
	PlayerID       int64  `json:"PlayerID"`
	GameToken      string `json:"GameToken"` // platform GameToken
	ExchangeType   int64  `json:"ExchangeType"`
	ExchangeAmount int64  `json:"ExchangeAmount"`
	TotalBet       int64  `json:"TotalBet"`
	TotalWin       int64  `json:"TotalWin"`
	TotalLost      int64  `json:"TotalLost"`
	IsCheckOut     bool   `json:"CheckOut"`

	AccountToken string `json:"AccountToken"` // Maintan checkout use
}

// Result plant ULG API Result
type Result struct {
	Result        int           `json:"result,omitempty"`
	AccountID     int64         `json:"userID,omitempty"`
	Status        int           `json:"status,omitempty"` // 0: empty 1:exchange 2:checkout
	AccountName   string        `json:"accountName,omitempty"`
	ErrorMsg      string        `json:"error_msg,omitempty"`
	UserName      string        `json:"userName,omitempty"`   // not use, give default value
	AccountToken  string        `json:"token,omitempty"`      // for plant token
	GameToken     string        `json:"game_token,omitempty"` // for game token
	UserPhone     string        `json:"userPhone,omitempty"`
	GameCoin      int64         `json:"gameCoin,omitempty"`
	UserCoinQuota []CoinQuota   `json:"userCoinQuota,CoinQuota,omitempty"`
	Coinsetting   []CoinSetting `json:"coinsetting,CoinSetting,omitempty"`
	GameInfo      []CoinInfo    `json:"gameInfo,CoinInfo,omitempty"`
	// CheckOutCoin  AmountCoin    `json:"amountCoin,AmountCoin,omitempty"`
}

// CheckOutResult Ulg check result
type CheckOutResult struct {
	Result        int       `json:"result"`
	ErrorMsg      string    `json:"errorMsg"`
	UserCoinQuota CoinQuota `json:"userCoinQuota,CoinQuota"`
}

// PartyAccount ...
func (ulg *Result) PartyAccount() string {
	return foundation.NewAccount("ulg", strconv.FormatInt(ulg.AccountID, 10))
}

// GameAccount ...
func (ulg *Result) GameAccount(encodeStr string) string {
	return foundation.NewGameAccount(encodeStr, string(ulg.AccountID))
}

// PartyToken ...
func (ulg *Result) PartyToken() string {
	return ulg.AccountToken
}

// AccountType ...
func (ulg *Result) AccountType() int64 {
	return playerinfo.Ulg
}

// CoinInfo Coin rate info
type CoinInfo struct {
	CoinType string  `json:"type"`
	Status   int     `json:"status"`
	Rate     float32 `json:"rate"`
	Sort     int     `json:"sort"`
}

// CoinQuota ulg user CoinQuota
type CoinQuota struct {
	CoinType string `json:"type,omitempty"`
	Amount   int64  `json:"amount"`

	Coin1Out     int64 `json:"coin1_out"`
	Coin2Out     int64 `json:"coin2_out"`
	Coin3Out     int64 `json:"coin3_out"`
	Coin4Out     int64 `json:"coin4_out"`
	Betting      int64 `json:"betting"`
	Win          int64 `json:"win,omitempty"`
	Lost         int64 `json:"lost,omitempty"`
	OutboundTime int64 `json:"outbound_time,omitempty"`
	Status       int   `json:"status,omitempty"`
}

// ToJSONClient ...
func (c CoinQuota) ToJSONClient() map[string]interface{} {
	result := make(map[string]interface{})
	result["cointype"] = c.CoinType
	result["amount"] = c.Amount
	return result
}

// CoinSetting ulg CoinSetting
type CoinSetting struct {
	Cointype string  `json:"cointype"` // money type
	Status   int     `json:"status"`   // enable status
	Rate     float32 `json:"rate"`     // exchange rate
	Sort     int     `json:"sort"`     // sort index

}

// AmountCoin coin check out
type AmountCoin struct {
	Coin1 int64 `json:"coin1"`
	Coin2 int64 `json:"coin2"`
	Coin3 int64 `json:"coin3"`
	Coin4 int64 `json:"coin4"`
}

// platform api url
var (
	ULGMaintainCheckoutTime = ""
	LoginURL                = ""
	GetuserURL              = ""
	AuthorizedURL           = ""
	ExchangeURL             = ""
	CheckoutURL             = ""
)
