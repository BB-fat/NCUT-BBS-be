package model

import (
	qaPB "ncutbbs/proto/qa"
	"strings"
)

type Question struct {
	ID         uint
	Title      string
	AuthorID   uint
	CreateTime int64 `gorm:"autoCreateTime"`
	Views      int
	Content    string
	Pictures   string
}

func (q *Question) ToData() *qaPB.QuestionData {
	data := qaPB.QuestionData{
		Id:         int32(q.ID),
		AuthorId:   int32(q.AuthorID),
		CreateTime: q.CreateTime,
		Title:      q.Title,
		Views:      int32(q.Views),
		Content:    q.Content,
	}
	if len(q.Pictures) == 0 {
		data.Pictures = []string{}
	} else {
		data.Pictures = strings.Split(q.Pictures, ",")
	}
	return &data
}

type Answer struct {
	ID         uint
	AuthorID   uint
	QuestionID uint
	CreateTime int64 `gorm:"autoCreateTime"`
	Content    string
}

func (a *Answer) ToData(userID uint) *qaPB.AnswerData {
	data := qaPB.AnswerData{
		Id:         int32(a.ID),
		QuestionId: int32(a.QuestionID),
		CreateTime: a.CreateTime,
		Content:    a.Content,
	}
	user := User{
		ID: a.AuthorID,
	}
	DB.First(&user)
	data.Author = user.ToSimpleData()
	var al []AnswerLike
	res := DB.Where("answer_id = ?", a.ID).Find(&al)
	data.Likes = int32(res.RowsAffected)
	res = DB.Where("answer_id = ? AND user_id = ?", a.ID, userID).Find(&al)
	if res.RowsAffected > 0 {
		data.IsLike = true
	} else {
		data.IsLike = false
	}
	return &data
}

type AnswerLike struct {
	AnswerID uint
	UserID   uint
}
