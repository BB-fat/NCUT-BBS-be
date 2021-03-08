package server

import (
	"github.com/gin-gonic/gin"
	"ncutbbs/server/controller"
)

func setUrl(r *gin.Engine) {
	r.GET("/", controller.Test)
}
