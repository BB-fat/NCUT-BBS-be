package controller

import (
	"github.com/gin-gonic/gin"
	"ncutbbs/module/session"
	"ncutbbs/ncut_bbs_api_model/auth"
	"ncutbbs/server/controller/tool"
	"net/http"
)

func Test1(c *gin.Context) {
	req := &auth.LoginRequest{}
	tool.LoadMessage(c, req)
	c.String(http.StatusOK, tool.DumpMessage(&auth.LoginReply{
		Success: true,
		Message: req.AccountName,
	}))
}

func Test2(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func Test3(c *gin.Context) {
	s := session.Get(c.GetHeader("token"))
	session.RefreshToken(s)
}
