package attach

import "time"

// Setting ...
type Setting struct {
	UserID string
}

// IAttach ...
type IAttach interface {
	LoadData()
	Get(attachkind int64, attachtype int64) *Info
	SetValue(attachKind, attachType int64, SValue string, IValue int64)
	SetAttach(attach *Info)
	Save()
	// Clear()
}

// Info attach data
type Info struct {
	Kind       int64  `json:"kind"`
	Types      int64  `json:"types"`
	SValue     string `json:"sValue"`
	IValue     int64  `json:"iValue"`
	UpdateTime int64  `json:"updateTime"`
	IsDirty    bool
	IsDBData   bool `json:"isDBData"`
}

// GetKind ...
func (i *Info) GetKind() int64 {
	return i.Kind
}

// GetTypes ...
func (i *Info) GetTypes() int64 {
	return i.Types
}

// GetSValue ...
func (i *Info) GetSValue() string { return i.SValue }

// SetSValue ...
func (i *Info) SetSValue(value string) {
	i.IsDirty = true
	i.UpdateTime = time.Now().Unix()
	i.SValue = value
}

// GetIValue ...
func (i *Info) GetIValue() int64 { return i.IValue }

// SetIValue ...
func (i *Info) SetIValue(value int64) {
	i.IsDirty = true
	i.UpdateTime = time.Now().Unix()
	i.IValue = value
}

// GetUpdateTime ...
func (i *Info) GetUpdateTime() int64 { return i.UpdateTime }

// GetIsDirty ...
func (i *Info) GetIsDirty() bool { return i.IsDirty }

// GetIsDBData ...
func (i *Info) GetIsDBData() bool { return i.IsDBData }
