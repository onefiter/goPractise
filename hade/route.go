package main

import (
	"github.com/goPractise/hade/framework/gin"
	"github.com/goPractise/hade/framework/middleware"
)

func registerRouter(core *gin.Engine) {
	// static Route And HTTP Method
	// core.Get("/user/login", UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	subjectApi.Use(middleware.Test3())
	{
		// 动态路由
		subjectApi.DELETE("/:id", SubjectDelController)
		subjectApi.PUT("/:id", SubjectUpdateController)
		subjectApi.GET("/:id", SubjectListController)
		subjectApi.GET("/list/all", SubjectListController)

	}
}
