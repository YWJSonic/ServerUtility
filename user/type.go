package user

import (
	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/attach"
	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/playerinfo"
)

// Info ...
type Info struct {
	UserServerInfo *playerinfo.AccountInfo
	UserGameInfo   *playerinfo.Info
	IAttach        attach.IAttach
}

// LoadAttach Init attach data
func (i *Info) LoadAttach() {
	i.IAttach.LoadData()
}
