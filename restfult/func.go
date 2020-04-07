package restfult

// NewRestfultService ...
func NewRestfultService() *Service {
	return &Service{
		proxyData: make(map[string]Setting),
	}
}
