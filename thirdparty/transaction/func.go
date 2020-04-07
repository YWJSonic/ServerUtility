package transaction

import "github.com/YWJSonic/ServerUtility/myhttp"

// NewTransaction ...
func NewTransaction(path string, httpConn *myhttp.Service) *Service {
	return &Service{
		Path:     path,
		httpConn: httpConn,
	}
}
