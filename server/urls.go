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
		}
	}

	account := r.Group("/account")
	{
		account.POST("/login", controller.LoginByPassword)
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
