package account

import (
	"ncutbbs/model"
	"ncutbbs/module/session"
	accountPB "ncutbbs/proto/account"
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

func CreateAccount(accountName, password, realName string, sex int32, college string, grade int32, avatar string) (success bool, token, message string) {
	u := model.User{
		AccountName:   accountName,
		Password:      password,
		RealName:      realName,
		Sex:           int(sex),
		College:       college,
		AccountStatus: int(accountPB.AccountStatus_UNNAMED),
		Grade:         int(grade),
		Avatar:        avatar,
	}
	count := int64(0)
	model.DB.Where("account_name = ?", accountName).Model(u).Count(&count)
	if count > 0 {
		return false, "", "用户名重复，请换一个"
	}
	res := model.DB.Create(&u)
	if res.Error != nil {
		return false, "", "未知错误，请重试"
	}
	s := session.Create(u)
	return true, s.Token, ""
}

func CreateVerifyInfo(userID uint, image, remark string) (success bool, message string) {
	u := model.User{}
	model.DB.First(&u, userID)
	if u.ID == 0 {
		return false, "没有找到用户"
	}
	info := model.VerifyInfo{}
	model.DB.Where("user_id = ?", u.ID).First(info)
	if info.ID != 0 {
		return false, "不能重复提交实名申请"
	}
	info = model.VerifyInfo{
		UserID: u.ID,
		Image:  image,
		Remark: remark,
		Pass:   0,
	}
	model.DB.Create(&info)
	return true, ""
}
