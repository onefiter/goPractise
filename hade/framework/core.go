package framework

import "net/http"

// framework core
type Core struct {
}

// init framework
func NewCore() *Core {
	return &Core{}
}

// framework handler

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// TODO
}
