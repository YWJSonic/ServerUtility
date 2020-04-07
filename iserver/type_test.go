package iserver

import (
	"fmt"
	"testing"

	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/foundation"
)

func TestService_Launch(t *testing.T) {
	jsStr := `{
		"TransferURL":"34.94.53.107",
		"IP":"127.0.0.1",
		"PORT":"8000",
		"SocketPORT":"9000",
		"DBIP":"127.0.0.1",
		"DBPORT":"3306",
		"DBUser":"serverConnect",
		"DBPassword":"123qweasdzxc",
		"AccountEncodeStr":"yrgb$",
		"RedisURL":"127.0.0.1:6379",
		"MaintainStartTime":"10 16 * * *",
		"MaintainFinishTime": "15 16 * * *",
		"ULGMaintainCheckoutTime":"12 16 * * *",
		"Maintain":false,
		"ServerMod":true
	}`
	config := foundation.StringToJSON(jsStr)
	baseSetting := NewSetting()
	baseSetting.SetData(config)

	type args struct {
		setting Setting
	}
	tests := []struct {
		name string
		s    *Service
		args args
	}{
		struct {
			name string
			s    *Service
			args args
		}{
			s: NewService(),
			args: args{
				setting: baseSetting,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Launch(tt.args.setting)

			userproto, _, err := tt.s.Transfer.AuthUser("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1ODQ1NDI2NDgsImlzcyI6IjhkNzAwODUyLTliMGEtNDA5Zi1hZjFkLWE4NDFkNWNhN2Y0MCIsImp0aSI6IjkwYTFlZjIxLTNhMzMtNDdhZi05YTU5LTY4ZDI1MjA3YWU4MyJ9.jpJyVtOPv-_njWdtm7n6e-Sv-VecfdeE150_JTp32Cw")

			fmt.Println("---body---", userproto)
			fmt.Println("---err---", err)
		})
	}
}
