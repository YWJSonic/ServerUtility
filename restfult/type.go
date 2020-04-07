package restfult

import (
	"net/http"

	"github.com/YWJSonic/ServerUtility/httprouter"
	"github.com/YWJSonic/ServerUtility/messagehandle"
	"github.com/YWJSonic/ServerUtility/myhttp"
)

// // IRestfult ...
// type IRestfult interface {
// 	ConnectPool() *http.Client
// 	HTTPLisentRun(ListenIP string, HandleURL ...[]myhttp.RESTfulURL) (err error)
// 	ListenProxy(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
// }
// ConnType ...
const (
	Client  = "cli"
	Backend = "back"
)

// Setting ...
type Setting struct {
	RequestType string
	URL         string
	Fun         httprouter.Handle
	ConnType    string
}

// Service IRestfult
type Service struct {
	proxyData map[string]Setting
}

// HTTPLisentRun ...
func (restfult *Service) HTTPLisentRun(adders string, HandleURL ...[]Setting) (err error) {
	router := httprouter.New()

	for _, RESTfulURLArray := range HandleURL {
		for _, RESTfulURLvalue := range RESTfulURLArray {
			messagehandle.LogPrintf("Restfult Listen %v %s\n", RESTfulURLvalue.RequestType, RESTfulURLvalue.URL)

			restfult.proxyData[RESTfulURLvalue.URL] = RESTfulURLvalue
			if RESTfulURLvalue.RequestType == "GET" {
				router.GET("/"+RESTfulURLvalue.URL, RESTfulURLvalue.Fun)
			} else if RESTfulURLvalue.RequestType == "POST" {
				router.POST("/"+RESTfulURLvalue.URL, restfult.ListenProxy)
			}
			router.OPTIONS("/"+RESTfulURLvalue.URL, myhttp.Option)

		}
	}

	messagehandle.LogPrintln("Restfult run on", adders)

	err = http.ListenAndServe(adders, router)
	if err != nil {
		messagehandle.ErrorLogPrintln("ListenAndServe", err)
		return err
	}
	return nil
}

// ListenProxy ...
func (restfult *Service) ListenProxy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	RESTfulInfo := restfult.proxyData[r.RequestURI[1:]]
	myhttp.AddHeader(&w)
	switch RESTfulInfo.ConnType {
	case Client:
		RESTfulInfo.Fun(w, r, ps)
	case Backend:
		RESTfulInfo.Fun(w, r, ps)
	}
}
