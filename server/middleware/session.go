package middleware

import (
	"github.com/gin-gonic/gin"
	"ncutbbs/module/session"
	"ncutbbs/server/controller/tool"
	"net/http"
)

func Session() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := tool.GetToken(c)
		if len(token) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		s := session.Get(token)
		if s == nil || !session.Valid(s) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
