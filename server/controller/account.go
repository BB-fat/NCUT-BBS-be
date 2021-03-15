package controller

import (
	"github.com/gin-gonic/gin"
	"ncutbbs/module/account"
	accountPB "ncutbbs/proto/account"
	"ncutbbs/server/controller/tool"
	"net/http"
)

func LoginByPassword(c *gin.Context) {
	req := accountPB.LoginRequest{}
	tool.LoadMessage(c, &req)
	success, token, message := account.LoginByPassword(req.AccountName, req.Password)
	c.String(http.StatusOK, tool.DumpMessage(&accountPB.LoginReply{
		Success: success,
		Token:   token,
		Message: message,
	}))
}
