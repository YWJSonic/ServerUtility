package iserver

import (
	"sync"

	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/dbservice"
)

// NewSetting ..
func NewSetting() Setting {
	return Setting{
		mu: new(sync.RWMutex),
	}
}

// NewService ...
func NewService() *Service {
	return &Service{
		ShotDown: make(chan bool),
		DBs:      make(map[string]*dbservice.DB),
	}
}
