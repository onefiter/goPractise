package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	go func() {

		server.ListenAndServe()
	}()

	// 等待的 Goroutine 等待信号量、
	quit := make(chan os.Signal)
	// 监控信号：SIGINT， SIGTERM，SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// 这里会阻塞当前 Goroutine等待信号
	<-quit

	// 调用Server.Shutdown graceful结束
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
