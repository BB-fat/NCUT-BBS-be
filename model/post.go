package model

import (
	forumPB "ncutbbs/proto/forum"
	"strings"
)

/*
create table ncut_bbs.posts
(
	id int auto_increment
		primary key,
	author_id int null,
	create_time int null,
	update_time int null,
	title text null,
	views int null,
	likes int null,
	content text null,
	unlikes int null,
	pictures text null,
	constraint posts_users_id_fk
		foreign key (author_id) references ncut_bbs.users (id)
);
*/

type Post struct {
	ID         uint
	AuthorID   uint
	CreateTime int64 `gorm:"autoCreateTime"`
	UpdateTime int64 `gorm:"autoUpdateTime"`
	Title      string
	Views      int
	Content    string
	Pictures   string
}

func (p *Post) ToData(userID uint) *forumPB.PostData {
	data := forumPB.PostData{
		Id:         int32(p.ID),
		AuthorId:   int32(p.AuthorID),
		CreateTime: p.CreateTime,
		UpdateTime: p.UpdateTime,
		Title:      p.Title,
		Views:      int32(p.Views),
		Content:    p.Content,
	}
	if len(p.Pictures) == 0 {
		data.Pictures = []string{}
	} else {
		data.Pictures = strings.Split(p.Pictures, ",")
	}
	var like PostLike
	res := DB.Where("post_id = ?", p.ID).Find(&like)
	data.Likes = int32(res.RowsAffected)
	res = DB.Where("user_id = ? AND post_id = ?", userID, p.ID).Find(&like)
	if res.RowsAffected > 0 {
		data.IsLike = true
	} else {
		data.IsLike = false
	}
	return &data
}

type PostLike struct {
	UserID uint
	PostID uint
}

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
