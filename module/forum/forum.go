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

func GetOne(userID uint, postID uint) *forumPB.PostData {
	post := model.Post{
		ID: postID,
	}
	model.DB.Find(&post)
	return Model2Data(post, userID)
}

func GetAll(userID uint) []*forumPB.PostData {
	var posts []model.Post
	model.DB.Find(&posts)
	var data []*forumPB.PostData
	for i := 0; i < len(posts); i++ {
		data = append(data, Model2Data(posts[i], userID))
	}
	return data
}

func Model2Data(post model.Post, userID uint) *forumPB.PostData {
	data := forumPB.PostData{
		Id:         int32(post.ID),
		AuthorId:   int32(post.AuthorID),
		CreateTime: post.CreateTime,
		UpdateTime: post.UpdateTime,
		Title:      post.Title,
		Views:      int32(post.Views),
		Content:    post.Content,
	}
	if len(post.Pictures) == 0 {
		data.Pictures = []string{}
	} else {
		data.Pictures = strings.Split(post.Pictures, ",")
	}
	var like model.PostLike
	res := model.DB.Where("post_id = ?", post.ID).Find(&like)
	data.Likes = int32(res.RowsAffected)
	res = model.DB.Where("user_id = ? AND post_id = ?", userID, post.ID).Find(&like)
	if res.RowsAffected > 0 {
		data.IsLike = true
	} else {
		data.IsLike = false
	}
	return &data
}

func Like(userID, postID uint) bool {
	data := model.PostLike{}
	res := model.DB.Where("user_id = ? AND post_id = ?", userID, postID).Find(&data)
	if res.RowsAffected > 0 {
		return false
	}
	data.UserID = userID
	data.PostID = postID
	model.DB.Create(&data)
	return true
}

func UnLike(userID, postID uint) {
	data := model.PostLike{}
	res := model.DB.Where("user_id = ? AND post_id = ?", userID, postID).Find(&data)
	if res.RowsAffected == 0 {
		return
	}
	model.DB.Where("user_id = ? AND post_id = ?", userID, postID).Delete(&data)
}
