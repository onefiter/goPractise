package framework

import (
	"log"
	"net/http"
	"strings"
)

// framework core
type Core struct {
	router map[string]map[string]ControllerHandler // 二级map
}

// init framework
func NewCore() *Core {
	// 定义二级map
	getRouter := map[string]ControllerHandler{}
	postRouter := map[string]ControllerHandler{}
	putRouter := map[string]ControllerHandler{}
	deleteRouter := map[string]ControllerHandler{}

	// 将二级map写入一级map
	router := map[string]map[string]ControllerHandler{}

	router["GET"] = getRouter
	router["POST"] = postRouter
	router["PUT"] = putRouter
	router["DELETE"] = deleteRouter

	return &Core{router: router}
}

// 对应 Method = GET
func (c *Core) Get(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["GET"][upperUrl] = handler
}

// 对应 Method = POST
func (c *Core) POST(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["POST"][upperUrl] = handler
}

// 对应 Method = PUT
func (c *Core) Put(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["PUT"][upperUrl] = handler
}

// 对应 Method = DELETE
func (c *Core) Delete(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["POST"][upperUrl] = handler
}

// 匹配路由，如果没有匹配到， 返回nil
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	// uri和method 全部转换为大写，保证大小写不敏感

	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)
	upperUri := strings.ToUpper(uri)

	// 查找第一层map

	if methodHandlers, ok := c.router[upperMethod]; ok {
		// 查找第二层map
		if handler, ok := methodHandlers[upperUri]; ok {
			return handler
		}
	}

	return nil
}

// framework handler

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	log.Println("core.serveHTTP")
	// 封装自定义context
	ctx := NewContext(request, response)

	router := c.FindRouteByRequest(request)
	if router == nil {
		ctx.Json(404, "not found")
		return
	}

	// 调用路由函数，如果返回err，代表存在内部错误，返回500状态码
	if err := router(ctx); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}
