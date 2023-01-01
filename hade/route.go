package main

import (
	"time"

	"github.com/goPractise/hade/framework"
)

func registerRouter(core *framework.Core) {
	// static Route And HTTP Method
	// core.Get("/user/login", UserLoginController)
	core.Get("/user/login", framework.TimeoutHandler(UserLoginController, time.Second*30))

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectListController)
		subjectApi.Get("/list/all", SubjectListController)

	}
}
