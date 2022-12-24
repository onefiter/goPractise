package middleware

import (
	"fmt"
	"time"

	"github.com/blog-service/global"
	"github.com/blog-service/pkg/app"
	"github.com/blog-service/pkg/email"
	"github.com/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})

	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf(ctx, "panic recover err: %v", err)

				err := defailtMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间：%d", time.Now().Unix()),
					fmt.Sprintf("错误信息：%v", err),
				)

				if err != nil {
					global.Logger.Panicf("mail.SendMail err: %v", err)
				}

				app.NewResponse(ctx).ToErrorResponse(errcode.ServerError)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}

}
