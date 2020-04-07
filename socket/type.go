package socket

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/YWJSonic/ServerUtility/messagehandle"
	"github.com/gorilla/websocket"
)

// Message game service call back handle on socket read,
type Message struct {
	MessageType   int
	Message       []byte
	SocketConnKey string
}

// ConnInfo ...
type ConnInfo struct {
	mapKey  string
	c       *websocket.Conn
	service *Service
	// rx      <-chan []byte // out put message
	// tx         chan<- []byte // send message
	handleRead func(Message) error
}

// Close ...
func (ci *ConnInfo) Close() {

	ci.c.WriteControl(websocket.CloseMessage, []byte("asdf1"), time.Now().Add(time.Second))
	// ci.c.Close()
}

// Send ...
func (ci *ConnInfo) Send(mtype int, msg []byte) error {
	err := ci.c.WriteMessage(mtype, msg)
	if err != nil {
		log.Println("Socket Write:", err)
		return err
	}
	return nil
}

// Listen ...
func (ci *ConnInfo) Listen() {
	defer func() {
		ci.service.CloseConn(ci.mapKey)
	}()

	for {
		mtype, msg, err := ci.c.ReadMessage()
		fmt.Println("---   socket message   ---", mtype, msg, err, "---------")
		if err != nil {
			log.Printf("Socket Message %d, %s, %v", mtype, msg, err)
			break
		}

		if err := ci.handleRead(Message{mtype, msg, ci.mapKey}); err != nil {
			log.Printf("Socket Read: %s", err)
		}
	}
}

// Setting ...
type Setting struct {
	RequestType string
	URL         string
	Fun         func(http.ResponseWriter, *http.Request)
	ConnType    string
}

// Service ...
type Service struct {
	*websocket.Upgrader
	ConnMap map[string]*ConnInfo
}

// AddNewConn ...
func (s *Service) AddNewConn(mapKey string, newConn *websocket.Conn, readFunc func(Message) error) {
	if _, ok := s.ConnMap[mapKey]; ok {
		s.ConnMap[mapKey].Close()
	}
	ci := &ConnInfo{
		mapKey,
		newConn,
		s,
		readFunc,
	}
	s.ConnMap[mapKey] = ci
	go ci.Listen()
}

// CloseConn ...
func (s *Service) CloseConn(mapKey string) {
	s.ConnMap[mapKey].Close()
	s.removeConn(mapKey)
}
func (s *Service) removeConn(mapKey string) {
	delete(s.ConnMap, mapKey)
}

// HTTPLisentRun ...
func (s *Service) HTTPLisentRun(adders string, HandleURL ...[]Setting) (err error) {

	for _, SocketURLArray := range HandleURL {
		for _, SocketURLvalue := range SocketURLArray {
			messagehandle.LogPrintf("Socket Listen %s\n", SocketURLvalue.URL)
			http.HandleFunc("/"+SocketURLvalue.URL, SocketURLvalue.Fun)
		}
	}

	messagehandle.LogPrintln("Socket run on", adders)
	err = http.ListenAndServe(adders, nil)
	if err != nil {
		messagehandle.ErrorLogPrintln("ListenAndServe", err)
		return err
	}
	return nil
}
