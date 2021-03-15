package controller

import (
	"github.com/gin-gonic/gin"
	"ncutbbs/module/session"
	sessionPB "ncutbbs/proto/session"
	"ncutbbs/server/controller/tool"
	"net/http"
)

func RefreshToken(c *gin.Context) {
	s := session.Get(c.GetHeader("token"))
	if s == nil {
		c.String(http.StatusOK, tool.DumpMessage(&sessionPB.RefreshTokenReply{
			Success: false,
		}))
	}
	session.RefreshToken(s)
	c.String(http.StatusOK, tool.DumpMessage(&sessionPB.RefreshTokenReply{
		Success: true,
		Token:   s.Token,
	}))
}
