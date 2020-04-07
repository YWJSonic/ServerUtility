package transaction

import "gitlab.fbk168.com/gamedevjp/backend-utility/utility/myhttp"

// NewTransaction ...
func NewTransaction(path string, httpConn *myhttp.Service) *Service {
	return &Service{
		Path:     path,
		httpConn: httpConn,
	}
}
