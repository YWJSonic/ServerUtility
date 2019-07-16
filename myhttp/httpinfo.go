package myhttp

import (
	"gitlab.com/ServerUtility/httprouter"
)

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
