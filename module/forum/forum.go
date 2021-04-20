package forum

import (
	"ncutbbs/model"
	forumPB "ncutbbs/proto/forum"
	"strings"
)

func Create(authorID uint, title, content string, pictures string) model.Post {
	post := model.Post{
		AuthorID: authorID,
		Title:    title,
		Views:    1,
		Content:  content,
		Pictures: pictures,
	}
	model.DB.Create(&post)
	return post
}

func Delete(ID uint) bool {
	post := model.Post{
		ID: ID,
	}
	res := model.DB.Delete(&post)
	if res.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

func GetAll() []*forumPB.PostData {
	var posts []model.Post
	model.DB.Find(&posts)
	var data []*forumPB.PostData
	for i := 0; i < len(posts); i++ {
		data = append(data, Model2Data(posts[i]))
	}
	return data
}

func Model2Data(post model.Post) *forumPB.PostData {
	return &forumPB.PostData{
		Id:         int32(post.ID),
		AuthorId:   int32(post.AuthorID),
		CreateTime: post.CreateTime,
		UpdateTime: post.UpdateTime,
		Title:      post.Title,
		Views:      int32(post.Views),
		Likes:      int32(post.Likes),
		Content:    post.Content,
		Unlikes:    int32(post.Unlikes),
		Pictures:   strings.Split(post.Pictures, ","),
	}
}
