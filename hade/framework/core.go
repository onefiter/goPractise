package framework

import (
	"log"
	"net/http"
)

// framework core
type Core struct {
	router map[string]ControllerHandler
}

// init framework
func NewCore() *Core {
	return &Core{router: map[string]ControllerHandler{}}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

// framework handler

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// TODO
	log.Println("core.serveHTTP")
	ctx := NewContext(request, response)

	router := c.router["foo"]
	if router == nil {
		return
	}

	log.Println("core.router")
	router(ctx)
}
