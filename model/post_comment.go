package model

import forumPB "ncutbbs/proto/forum"

type PostComment struct {
	ID         uint
	AuthorID   uint
	PostID     uint
	CreateTime int64 `gorm:"autoCreateTime"`
	Content    string
}

func (p *PostComment) ToData() *forumPB.PostCommentData {
	data := forumPB.PostCommentData{
		Id:         int32(p.ID),
		PostId:     int32(p.ID),
		CreateTime: p.CreateTime,
		Content:    p.Content,
	}
	user := User{}
	DB.First(&user, p.AuthorID)
	data.Author = user.ToSimpleData()
	return &data
}
