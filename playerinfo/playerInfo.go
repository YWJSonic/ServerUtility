package playerinfo

// Info Player information
type Info struct {
	IDStr string `json:"IDStr"`
	ID    int64  `json:"ID"`
	// Money       int64  `json:"Money"`
	MoneyU      uint64 `json:"MoneyU"`
	GameAccount string `json:"GameAccount"`
	Name        string `json:"Name"`

	///////// for Server value
	GameToken     string `json:"GameToken,omitempty"`
	InRoom        int    `json:"InRoom,omitempty"`        // room index
	LastCheckTime int64  `json:"LastCheckTime,omitempty"` // connect check time
	InGame        string `json:"InGame,omitempty"`        // gametype
}

// GetMoney ...
func (p *Info) GetMoney() int64 {
	return int64(p.MoneyU)
}

// GetMoneyU ...
func (p *Info) GetMoneyU() uint64 {
	return p.MoneyU
}

// SumMoney ...
func (p *Info) SumMoney(value int64) int64 {
	p.MoneyU += uint64(value)
	return p.GetMoney()
}

// SetMoney ...
func (p *Info) SetMoney(value int64) {
	p.MoneyU = uint64(value)
}

// ToJSONClient ...
func (p Info) ToJSONClient() map[string]interface{} {
	clientdata := make(map[string]interface{})
	clientdata["id"] = p.ID
	clientdata["Money"] = p.GetMoney()
	clientdata["gameaccount"] = p.GameAccount
	return clientdata
}

// ResultMap player conver to map, client data
func (p Info) ResultMap() map[string]interface{} {
	return map[string]interface{}{
		"ID":    p.ID,
		"Money": p.GetMoney(),
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
