package main

import (
	"net/http"

	"github.com/goPractise/hade/framework"
)

func main() {
	server := &http.Server{
		Handler: framework.NewCore(),
		Addr:    "localhost:8080",
	}
	server.ListenAndServe()
}
