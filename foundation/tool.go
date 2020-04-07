package foundation

import (
	"fmt"
	"reflect"
	"runtime"
	"time"

	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/code"
	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/messagehandle"
)

// ServerTotalPayScoreKey ...
func ServerTotalPayScoreKey(GameIndex int) string {
	return fmt.Sprintf("ServerTotalPayScore%d", GameIndex)
}

// NewAccount convert all plant account to server account
func NewAccount(plant, account string) string {
	return fmt.Sprintf("%s:%s", plant, account)
}

// NewGameAccount new game account
func NewGameAccount(encodeStr, account string) string {
	return MD5Code(encodeStr + account)
}

// NewToken ...
func NewToken(gameAccount string) string {
	return MD5Code(fmt.Sprintf("%s%d", gameAccount, ServerNowTime()))
}

// CheckToken Check Token func
func CheckToken(serverToken, clientToken string) messagehandle.ErrorMsg {
	err := messagehandle.New()
	if serverToken != clientToken {
		err.ErrorCode = code.Unauthenticated
		err.Msg = "TokenError"
	}
	return err
}

// CheckGameType Check Game Type
func CheckGameType(serverGameTypeID, clientGameTypeID string) messagehandle.ErrorMsg {
	err := messagehandle.New()
	if serverGameTypeID != clientGameTypeID {
		err.ErrorCode = code.GameTypeError
		err.Msg = "GameTypeError"
	}
	return err
}

// ServerNowTime Get now Unix time
func ServerNowTime() int64 {
	return time.Now().Unix()
}

// ServerNow Get now time
func ServerNow() time.Time {
	return time.Now()
}

// IsInclude ...
func IsInclude(target int, src []int) bool {
	for _, value := range src {
		if value == target {
			return true
		}
	}
	return false
}

// GetFuncName return func Name
func GetFuncName(fun interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fun).Pointer()).Name()
}
