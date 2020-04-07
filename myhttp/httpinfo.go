package myhttp

import "github.com/YWJSonic/ServerUtility/httprouter"

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
