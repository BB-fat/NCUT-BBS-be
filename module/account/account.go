package account

import (
	"ncutbbs/model"
	"ncutbbs/module/session"
)

func LoginByPassword(accountName, password string) (success bool, token string, message string) {
	u := model.User{}
	model.DB.Where("account_name = ? AND password = ?", accountName, password).First(&u)
	if u.ID == 0 {
		return false, "", "用户名或密码错误"
	}
	s := model.Session{}
	model.DB.Where("user_id = ?", u.ID).First(&s)
	if s.ID == 0 {
		s = session.Create(u)
	} else {
		session.RefreshToken(&s)
	}
	return true, s.Token, ""
}
