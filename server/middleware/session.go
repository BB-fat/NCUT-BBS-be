package middleware

import (
	"github.com/gin-gonic/gin"
	"ncutbbs/module/session"
	"net/http"
)

func Session() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if len(token) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		s := session.Get(token)
		if s == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !session.Valid(s) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
