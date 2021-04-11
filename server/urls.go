package server

import (
	"github.com/gin-gonic/gin"
	"ncutbbs/config"
	"ncutbbs/server/controller"
	"ncutbbs/server/middleware"
)

func setUrl(r *gin.Engine) {
	r.Static("/static", "./static")

	v1 := r.Group("/v1")
	{
		// 刷新Token
		v1.GET("/refresh-token", controller.RefreshToken)

		// 上传文件
		v1.POST("/upload", controller.Upload)

		account := v1.Group("/account")
		{
			account.POST("/login", controller.LoginByPassword)
			// 创建账号
			account.POST("/create", controller.CreateAccount)

			accountSession := account.Group("")
			accountSession.Use(middleware.Session())
			{
				accountSession.GET("/user-info", controller.GetUserInfo)

				accountSession.POST("/verify", controller.SubmitVerfiyInfo)
			}
		}
	}

	// 测试接口
	if config.Config.Debug {
		test := r.Group("text")
		{
			test.GET("/1", controller.Test1)
			test.GET("/3", controller.Test3)
		}
	}
}
