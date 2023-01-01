package main

import (
	"github.com/goPractise/hade/framework"
	"github.com/goPractise/hade/framework/middleware"
)

func registerRouter(core *framework.Core) {
	// static Route And HTTP Method
	// core.Get("/user/login", UserLoginController)

	core.Use(
		middleware.Test1(),
		middleware.Test2(),
	)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	subjectApi.Use(middleware.Test3())
	{
		// 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectListController)
		subjectApi.Get("/list/all", SubjectListController)

	}
}
