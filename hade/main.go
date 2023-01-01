package main

import (
	"net/http"

	"github.com/goPractise/hade/framework"
	"github.com/goPractise/hade/framework/middleware"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)

	core.Use(middleware.Recovery())
	server := &http.Server{
		Handler: core,
		Addr:    "localhost:8080",
	}
	server.ListenAndServe()
}
