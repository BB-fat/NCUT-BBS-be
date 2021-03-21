package server

import (
	"github.com/gin-gonic/gin"
	"ncutbbs/config"
	"ncutbbs/server/controller"
	"ncutbbs/server/middleware"
)

func setUrl(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		session := v1.Group("")
		session.Use(middleware.Session())
		{
			session.GET("/test2", controller.Test2)
			// 刷新Token
			session.GET("/refresh-token", controller.RefreshToken)
		}

		account := v1.Group("/account")
		{
			account.POST("/login", controller.LoginByPassword)

			accountSession := account.Group("")
			accountSession.Use(middleware.Session())
			{
				accountSession.GET("/user-info", controller.GetUserInfo)
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
