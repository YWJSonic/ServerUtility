package messagehandle

import (
	"github.com/YWJSonic/ServerUtility/code"
)

// ErrorMsg ...
type ErrorMsg struct {
	ErrorCode code.Code
	// MsgNum    int8
	Msg string
}

func (e *ErrorMsg) Error() string {
	return e.Msg
}

// New default Error Message
func New() ErrorMsg {
	return ErrorMsg{
		ErrorCode: code.OK,
		// MsgNum:    0,
		Msg: "",
	}
}
