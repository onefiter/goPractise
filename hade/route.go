package main

import "github.com/goPractise/hade/framework"

func registerRouter(core *framework.Core) {
	// static Route And HTTP Method
	core.Get("/user/login", UserLoginController)

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
