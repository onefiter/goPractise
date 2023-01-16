package main

import (
	"fmt"

	"github.com/goPractise/hade/framework/gin"
	"github.com/goPractise/hade/provider/demo"
)

func SubjectAddController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectAddController")

}

func SubjectListController(c *gin.Context) {
	// 服务实例化
	demoService := c.MustMake(demo.Key).(demo.Service)

	foo := demoService.GetFoo()

	c.ISetOkStatus().IJson(foo)

}

func SubjectDelController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectDelController")
}

func SubjectUpdateController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectUpdateController")

}

func SubjectGetController(c *gin.Context) {
	subjectId, _ := c.DefaultParamInt("id", 0)
	c.ISetOkStatus().IJson("ok, SubjectGetController:" + fmt.Sprint(subjectId))
}

func SubjectNameController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectNameController")
}
