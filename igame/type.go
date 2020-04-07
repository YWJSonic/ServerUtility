package igame

import (
	"github.com/YWJSonic/ServerUtility/attach"
	"github.com/YWJSonic/ServerUtility/restfult"
	"github.com/YWJSonic/ServerUtility/socket"
)

// IGame ...
type IGame interface {
	RESTfulURLs() []restfult.Setting
	SocketURLs() []socket.Setting
	// GetUser(userToken string) (*user.Info, error)
	// GetUserByGameID(userToken string, UserID int64) (*user.Info, error)
	CheckToken(userToken string) error
}

// IRule ...
type IRule interface {
	GetBetMoney(index int64) int64
	GetGameTypeID() string
	GameRequest(*RuleRequest) *RuleRespond
	CheckGameType(userGameTypeID string) bool
	GetBetSetting() map[string]interface{}
}

// ISlotRule ...
type ISlotRule interface {
	IRule
	GetReel() map[string][][]int
}

// RuleRequest game logic request params
type RuleRequest struct {
	Attach   *attach.IAttach // copy data
	BetIndex int64
	UserID   int64
}

// RuleRespond game logic respond params
type RuleRespond struct {
	Attach        []*attach.Info
	BetMoney      int64
	Totalwinscore int64
	GameResult    map[string]interface{}
	OtherData     map[string]interface{}
}
