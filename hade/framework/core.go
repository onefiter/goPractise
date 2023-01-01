package framework

import (
	"log"
	"net/http"
	"strings"
)

// framework core
type Core struct {
	router      map[string]*Tree
	middlewares []ControllerHandler
}

// init framework
func NewCore() *Core {
	// 初始化路由
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{router: router}
}

// 对应 Method = GET
func (c *Core) Get(url string, handlers []ControllerHandler) {

	// 将core的middleware和handlers结合起来
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}

}

// 对应 Method = POST
func (c *Core) Post(url string, handlers []ControllerHandler) {

	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// 对应 Method = PUT
func (c *Core) Put(url string, handlers []ControllerHandler) {

	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// 对应 Method = DELETE
func (c *Core) Delete(url string, handlers []ControllerHandler) {

	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

// 匹配路由，如果没有匹配到， 返回nil
func (c *Core) FindRouteByRequest(request *http.Request) []ControllerHandler {
	// uri和method 全部转换为大写，保证大小写不敏感
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	// 查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {

		return methodHandlers.FindHandler(uri)
	}

	return nil
}

// framework handler

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	// 封装自定义context
	ctx := NewContext(request, response)

	handlers := c.FindRouteByRequest(request)
	if handlers == nil {
		ctx.Json(404, "not found")
		return
	}

	ctx.SetHandlers(handlers)

	// 调用路由函数，如果返回err，代表存在内部错误，返回500状态码
	if err := ctx.Next(); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}

func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)

}
