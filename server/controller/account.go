package controller

import (
	"github.com/gin-gonic/gin"
	"ncutbbs/module/account"
	"ncutbbs/ncut_bbs_api_model/accountPB"
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
