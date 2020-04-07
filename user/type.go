package user

import (
	"github.com/YWJSonic/ServerUtility/attach"
	"github.com/YWJSonic/ServerUtility/playerinfo"
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
