package myhttp

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/YWJSonic/ServerUtility/httprouter"
	"github.com/YWJSonic/ServerUtility/messagehandle"
)

// PostData get http post data
func PostData(r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	contentType := r.Header.Get("Content-type")

	if contentType == "application/x-www-form-urlencoded" {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		v := r.Form
		postdata := v.Get("POST")
		if err := json.Unmarshal([]byte(postdata), &data); err != nil {
			panic(err)
		}

	} else {
		d := json.NewDecoder(r.Body)
		err := d.Decode(&data)
		if err != nil {
			panic(err)
		}
	}

	return data
}

// AddHeader add request header
func AddHeader(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	// (*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Content-Type", "application/json")
	// (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// Option add header option
func Option(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	headers := w.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "*")
	headers.Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.WriteHeader(http.StatusOK)
}

// Service ...
type Service struct {
	client *http.Client
}

// ConnectPool ...
func (s *Service) ConnectPool() *http.Client {
	if s.client == nil {
		s.client = new(http.Client)
		httptr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},

			MaxIdleConns:        50,
			MaxIdleConnsPerHost: 50,
		}
		s.client = &http.Client{
			Transport: httptr,
		}
	}
	return s.client
}

// GET ...
func (s *Service) GET(url string, heard map[string][]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header = heard
	res, err := s.ConnectPool().Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return body, errors.New(res.Status)
	}

	return body, nil
}

// POST Http Raw Request
func (s *Service) POST(url string, value []byte, header map[string][]string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(value))
	if _, ok := header["Content-Type"]; !ok {
		header["Content-Type"] = []string{"application/json"}
	}
	req.Header = header
	res, err := s.ConnectPool().Do(req)
	if err != nil {
		messagehandle.ErrorLogPrintln("HTTPPostRawRequest Res", res)
		messagehandle.ErrorLogPrintln("HTTPPostRawRequest Error", err)
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode/100 != 2 {
		return body, errors.New(res.Status)
	}

	return body, nil
}

// PUT Http Raw Request
func (s *Service) PUT(url string, value []byte, header map[string][]string) ([]byte, error) {
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(value))
	if _, ok := header["Content-Type"]; !ok {
		header["Content-Type"] = []string{"application/json"}
	}
	req.Header = header
	res, err := s.ConnectPool().Do(req)
	if err != nil {
		messagehandle.ErrorLogPrintln("HTTPPostRawRequest Res", res)
		messagehandle.ErrorLogPrintln("HTTPPostRawRequest Error", err)
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode/100 != 2 {
		return body, errors.New(res.Status)
	}

	return body, nil
}

// HTTPGet ...
func (s *Service) HTTPGet(url string, values map[string][]string) ([]byte, error) {
	res, err := s.ConnectPool().Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

// HTTPPostRequest ...
func (s *Service) HTTPPostRequest(url string, values map[string][]string) ([]byte, error) {
	res, err := s.ConnectPool().PostForm(url, values)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)

}

// HTTPPostRawRequest Http Raw Request
func (s *Service) HTTPPostRawRequest(url string, value []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(value))
	req.Header.Set("Content-Type", "application/json")
	messagehandle.LogPrintln("HTTPPostRawRequest", req)

	resp, err := s.ConnectPool().Do(req)
	defer resp.Body.Close()
	if err != nil {
		messagehandle.ErrorLogPrintln("HTTPPostRawRequest Resp", resp)
		messagehandle.ErrorLogPrintln("HTTPPostRawRequest Error", err)
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}
