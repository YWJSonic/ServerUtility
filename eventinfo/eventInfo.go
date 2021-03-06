package eventinfo

import "time"

// event type enum
const (
	LoopEvent = iota + 1
	AtTimeEvent
)

// Info UpdateEvent struct
// DoTime: run at server  time on time
// LoopTime: run at server every pass time.Duration time
// *LoopTime limit Mini time.Duration is 1 Second
type Info struct {
	EventType int
	CountTime int64
	DoTime    int64
	LoopTime  time.Duration
	Do        func(interface{}) interface{}
	IsLaunch  bool
}
