package transaction

import (
	"errors"

	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/myhttp"
)

// User ...
type User struct {
}

// Request ...
type Request struct {
	bet     int64
	win     int64
	UserID  string
	Token   string
	GameID  string
	OrderID string
}

// Reply ...
type Reply struct {
}

// Service ...
type Service struct {
	Path     string
	httpConn *myhttp.Service
}

// AuthUser GET transation
func (s *Service) AuthUser(url string) ([]byte, error) {
	res, err := s.httpConn.GET(url, map[string][]string{})
	if len(res) <= 0 {
		if err != nil {
			return nil, err
		}
		return nil, errors.New(url + " return empty data.")
	}

	if err != nil {
		return res, err
	}

	return res, nil
}

// NewOrder POST transation
func (s Service) NewOrder(url, token string, payload []byte) ([]byte, error) {
	header := map[string][]string{
		"Authorization": []string{token},
		"Content-Type":  []string{"application/protobuf"},
	}

	res, err := s.httpConn.POST(url, payload, header)
	if len(res) <= 0 {
		if err != nil {
			return nil, err
		}
		return nil, errors.New(url + " return empty data.")
	}
	if err != nil {
		return res, err
	}

	return res, nil
}

// EndOrder GET transation
func (s Service) EndOrder(url, token string, payload []byte) ([]byte, error) {
	header := map[string][]string{
		"Authorization": []string{token},
	}

	res, err := s.httpConn.PUT(url, payload, header)
	if len(res) <= 0 {
		if err != nil {
			return nil, err
		}
		return nil, errors.New(url + " return empty data.")
	}

	if err != nil {
		return res, err
	}
	return res, nil
}
