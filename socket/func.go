package socket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// NewSocket ...
func NewSocket() *Service {
	socket := Service{
		&websocket.Upgrader{
			//如果有 cross domain 的需求，可加入這個，不檢查 cross domain
			CheckOrigin: func(r *http.Request) bool { return true },
		},
		make(map[string]*ConnInfo),
	}
	return &socket
}
