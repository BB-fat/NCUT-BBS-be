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
		UserInfo: &accountPB.UserInfo{
			Id:            int64(u.ID),
			AccountName:   u.AccountName,
			RealName:      u.RealName,
			Sex:           accountPB.Sex(u.Sex),
			College:       u.College,
			AccountStatus: accountPB.AccountStatus(u.AccountStatus),
			Grade:         int32(u.Grade),
			Avatar:        u.Avatar,
		},
	}))
}

func CreateAccount(c *gin.Context) {
	req := accountPB.CreateAccountRequest{}
	tool.LoadMessage(c, &req)
	success, token, message := account.CreateAccount(req.AccountName, req.Password, req.RealName, int32(req.Sex), req.College, req.Grade, req.Avatar)
	c.String(http.StatusOK, tool.DumpMessage(&accountPB.CreateAccountReply{
		Success: success,
		Token:   token,
		Message: message,
	}))
}

func SubmitVerifyInfo(c *gin.Context) {
	req := accountPB.SubmitVerifyInfoRequest{}
	tool.LoadMessage(c, &req)
	u := tool.GetUser(c)
	success, message := account.CreateVerifyInfo(u.ID, req.Image, req.Remark)
	c.String(http.StatusOK, tool.DumpMessage(&accountPB.SubmitVerifyInfoReply{
		Success: success,
		Message: message,
	}))
}
