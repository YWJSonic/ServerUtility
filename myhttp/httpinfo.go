package myhttp

import "gitlab.fbk168.com/gamedevjp/backend-utility/utility/httprouter"

// ConnType ...
const (
	Client  = "cli"
	Backend = "back"
)

// RESTfulURL ...
type RESTfulURL struct {
	RequestType string
	URL         string
	Fun         httprouter.Handle
	ConnType    string
}
