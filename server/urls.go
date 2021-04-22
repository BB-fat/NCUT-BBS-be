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

				accountSession.POST("/verify", controller.SubmitVerifyInfo)
			}
		}

		// 论坛
		forum := v1.Group("/forum")
		forum.Use(middleware.Session())
		{
			forum.POST("/post", controller.CreatePost)
			forum.GET("/post", controller.GetPost)

			forum.GET("/all", controller.GetPostList)

			forum.POST("/like", controller.LikePost)
			forum.DELETE("/like", controller.UnLikePost)

			forum.POST("/views", controller.AddPostViews)
		}
	}

	// 测试接口
	if config.Config.Debug {
	}
}
