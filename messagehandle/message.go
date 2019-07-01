package messagehandle

import "gitlab.com/ServerUtility/code"

// ErrorMsg ...
type ErrorMsg struct {
	ErrorCode code.Code
	// MsgNum    int8
	Msg string
}

// New default Error Message
func New() ErrorMsg {
	return ErrorMsg{
		ErrorCode: code.OK,
		// MsgNum:    0,
		Msg: "",
	}
}
