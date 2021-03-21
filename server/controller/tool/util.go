package tool

import (
	"github.com/gin-gonic/gin"
	"ncutbbs/model"
	"ncutbbs/module/session"
)

func GetToken(c *gin.Context) string {
	return c.GetHeader("token")
}

func GetUser(c *gin.Context) *model.User {
	return session.GetUser(GetToken(c))
}
