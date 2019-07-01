package playerinfo

import (
	"time"
)

// Account Type
const (
	None = iota
	Guest
	Self
	Ulg
)

// IPratyAccount thirdparty api interface
type IPratyAccount interface {
	PartyAccount() string
	GameAccount() string
	AccountType() int64
}

// AccountInfo ...
type AccountInfo struct {
	Account     string `json:"Account"`
	GameAccount string `json:"GameAccount"`
	AccountType int64  `json:"AccountType"`
	LoginTime   int64  `json:"LoginTime"`

	AccountToken string `json:"AccountToken"` // platform AccountToken
	Token        string `json:"Token"`        // Server Token
}

// NewAccountInfo ...
func NewAccountInfo(account, gameAccount, token string, accountType int64) AccountInfo {
	return AccountInfo{
		Account:     account,
		GameAccount: gameAccount,
		Token:       token,
		LoginTime:   time.Now().Unix(),
		AccountType: accountType,
	}
}
