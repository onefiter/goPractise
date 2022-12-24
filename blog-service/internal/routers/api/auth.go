package api

import (
	"log"

	"github.com/blog-service/global"
	"github.com/blog-service/internal/service"
	"github.com/blog-service/pkg/app"
	"github.com/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)

	log.Println("aaa", valid)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())

	err := svc.CheckAuth(&param)

	if err != nil {
		global.Logger.Errorf(c, "svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)

	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)

	}

	response.ToResponse(gin.H{
		"token": token,
	})

}
