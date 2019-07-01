package playerinfo

// Info Player information
type Info struct {
	ID          int64  `json:"ID"`
	Money       int64  `json:"Money"`
	GameAccount string `json:"GameAccount"`

	///////// for Server value
	GameToken     string `json:"GameToken,omitempty"`
	InRoom        int    `json:"InRoom,omitempty"`        // room index
	LastCheckTime int64  `json:"LastCheckTime,omitempty"` // connect check time
	InGame        string `json:"InGame,omitempty"`        // gametype
}

// ToJSONClient ...
func (p Info) ToJSONClient() map[string]interface{} {
	clientdata := make(map[string]interface{})
	clientdata["id"] = p.ID
	clientdata["money"] = p.Money
	clientdata["gameaccount"] = p.GameAccount
	return clientdata
}

// ResultMap player conver to map, client data
func (p Info) ResultMap() map[string]interface{} {
	return map[string]interface{}{
		"ID":    p.ID,
		"Money": p.Money,
		// "Token": p.Token,
	}
}

// IsInGameRoom is player in game room
func (p Info) IsInGameRoom() bool {
	return p.InRoom != 0
}

// IsCheckOut is ulg checkout
func (p Info) IsCheckOut() bool {
	return p.GameToken == ""
}
