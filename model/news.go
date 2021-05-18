package model

import newsPB "ncutbbs/proto/news"

type News struct {
	ID         uint
	CreateTime int64 `gorm:"autoCreateTime"`
	Type       int
	Title      string
	Content    string
	Data       string
	UserID     uint
	Received   int
}

func (n *News) ToData() *newsPB.NewsData {
	return &newsPB.NewsData{
		Id:         int32(n.ID),
		CreateTime: n.CreateTime,
		Type:       newsPB.NewsType(n.Type),
		Title:      n.Title,
		Content:    n.Content,
		Data:       n.Data,
	}
}
