package session

import (
	uuid "github.com/satori/go.uuid"
	"ncutbbs/config"
	"ncutbbs/model"
	"time"
)

func Get(token string) *model.Session {
	session := model.Session{}
	model.DB.Where("token = ?", token).First(&session)
	if session.ID == 0 {
		return nil
	} else {
		return &session
	}
}

func Create(user model.User) model.Session {
	session := model.Session{
		UserID: user.ID,
		Token:  uuid.NewV4().String(),
	}
	res := model.DB.Create(&session)
	if res.Error != nil {
		// TODO 创建失败
	}
	return session
}

func RefreshToken(session *model.Session) {
	model.DB.Model(session).Update("token", uuid.NewV4().String())
}

func Valid(session *model.Session) bool {
	return time.Now().Unix()-session.UpdateTime <= int64(config.Config.SessionTTL)
}

func GetUser(token string) *model.User {
	s := Get(token)
	if s == nil {
		return nil
	}
	u := model.User{}
	model.DB.Where("id = ?", s.UserID).First(&u)
	return &u
}
