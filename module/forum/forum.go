package forum

import (
	"fmt"
	"ncutbbs/model"
	"ncutbbs/module/news"
	forumPB "ncutbbs/proto/forum"
	newsPB "ncutbbs/proto/news"
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
	return post.ToData(userID)
}

func GetAll(userID uint) []*forumPB.PostData {
	var posts []model.Post
	model.DB.Find(&posts)
	var data []*forumPB.PostData
	for i := 0; i < len(posts); i++ {
		data = append(data, posts[i].ToData(userID))
	}
	return data
}

func Like(user *model.User, postID uint) bool {
	data := model.PostLike{}
	res := model.DB.Where("user_id = ? AND post_id = ?", user.ID, postID).Find(&data)
	if res.RowsAffected > 0 {
		return false
	}
	data.UserID = user.ID
	data.PostID = postID
	model.DB.Create(&data)

	post := model.Post{ID: postID}
	model.DB.First(&post)
	if post.AuthorID != user.ID {
		news.CreateNews(newsPB.NewsType_FORUM, fmt.Sprintf("%s赞了你的帖子", user.AccountName), "轻点查看详情", post.AuthorID, post.ToData(post.AuthorID))
	}

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

func AddViews(postID uint) {
	var post model.Post
	model.DB.First(&post, postID)
	post.Views++
	model.DB.Save(&post)
}

func CreateComment(user *model.User, postID uint, content string) *forumPB.PostCommentData {
	data := model.PostComment{
		AuthorID: user.ID,
		PostID:   postID,
		Content:  content,
	}
	model.DB.Create(&data)

	post := model.Post{ID: postID}
	model.DB.First(&post)
	if post.AuthorID != user.ID {
		news.CreateNews(newsPB.NewsType_FORUM, fmt.Sprintf("%s评论了你的帖子", user.AccountName), content, post.AuthorID, post.ToData(user.ID))
	}

	return data.ToData()
}

func GetPostComments(postID uint) []*forumPB.PostCommentData {
	var data []model.PostComment
	model.DB.Where("post_id = ?", postID).Find(&data)
	list := make([]*forumPB.PostCommentData, len(data), len(data))
	for i := 0; i < len(data); i++ {
		list[i] = data[i].ToData()
	}
	return list
}
