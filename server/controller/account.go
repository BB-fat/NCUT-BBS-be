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

func GetUserInfo(c *gin.Context) {
	u := tool.GetUser(c)
	c.String(http.StatusOK, tool.DumpMessage(&accountPB.GetUserInfoReply{
		Id:            int64(u.ID),
		AccountName:   u.AccountName,
		RealName:      u.RealName,
		Sex:           accountPB.GetUserInfoReply_Sex(u.Sex),
		College:       u.College,
		AccountStatus: accountPB.GetUserInfoReply_AccountStatus(u.AccountStatus),
		Grade:         int32(u.Grade),
		Avatar:        u.Avatar,
	}))
}
