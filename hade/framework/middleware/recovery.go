package middleware

import "github.com/goPractise/hade/framework"

// recovery机制，将协程中的函数异常进行捕获

func Recovery() framework.ControllerHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		defer func() {
			if err := recover(); err != nil {
				c.Json(500)
			}
		}()

		// 使用next执行具体的业务逻辑
		c.Next()
		return nil
	}

}
