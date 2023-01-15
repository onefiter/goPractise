package middleware

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/goPractise/hade/framework"
)

// 超时中间件

func TimeoutHandler(d time.Duration) framework.ControllerHandler {
	// 使用回调函数
	// 使用函数回调
	return func(c *framework.Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		//执行业务逻辑前预操作，初始化超时 context
		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			// 执行具体的业务逻辑
			c.Next()

			finish <- struct{}{}
		}()

		// 执行业务逻辑后操作
		select {
		case p := <-panicChan:
			log.Println(p)
			c.Json("timeout")
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.SetHasTimeout()
			c.Json("time out")
		}

		return nil

	}
}
