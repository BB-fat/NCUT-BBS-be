package controller

import (
	"github.com/gin-gonic/gin"
	"ncutbbs/ncut_bbs_api_model/auth"
	"net/http"
)

func Test1(c *gin.Context) {
	req := &auth.LoginRequest{}
	loadMessage(c, req)
	c.String(http.StatusOK, dumpMessage(&auth.LoginReply{
		Success: true,
		Message: req.AccountName,
	}))
}
