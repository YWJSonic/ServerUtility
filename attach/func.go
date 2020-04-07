package attach

// NewInfo ...
func NewInfo(kind, types int64, isDBData bool) *Info {
	return &Info{
		Kind:     kind,
		Types:    types,
		IsDBData: isDBData,
	}
}

// NewUserAttach ...
func NewUserAttach(attSetting Setting) *UserAttach {
	attach := &UserAttach{
		userID:  attSetting.UserID,
		dataMap: make(map[int64]map[int64]*Info),
	}
	// attach.InitData(userID)
	return attach
}
