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
	CreateTime int64 `gorm:"autoCreateTime"`
	Content    string
}

type AnswerLike struct {
	AnswerID uint
	UserID   uint
}
