package settinginfo

import (
	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/foundation"
)

// Info server setting info
type Info struct {
	Key         string
	IValue      int64
	SValue      string
	LastRefulsh int64
}

// ConvertToInfo db data to setting info
func ConvertToInfo(data map[string]interface{}) Info {
	var info Info

	info.Key = foundation.InterfaceToString(data["Key"])
	info.IValue = foundation.InterfaceToInt64(data["IValue"])
	info.SValue = foundation.InterfaceToString(data["SValue"])
	info.LastRefulsh = foundation.InterfaceToInt64(data["LastRefresh"])
	return info
}
