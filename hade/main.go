package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/goPractise/hade/framework"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)

	server := &http.Server{
		Handler: core,
		Addr:    "localhost:8080",
	}
	server.ListenAndServe()
}

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}

func FooControllerHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)
	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		time.Sleep(10 * time.Second)
		c.Json(200, "ok")
		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Panicln(p)
		c.Json(500, "panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()

		c.Json(500, "timeout")
		c.SetHasTimeout()
	}

	return nil
}
