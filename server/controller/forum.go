package controller

import (
	"github.com/gin-gonic/gin"
	"ncutbbs/module/forum"
	commonPB "ncutbbs/proto/common"
	forumPB "ncutbbs/proto/forum"
	"ncutbbs/server/controller/tool"
	"net/http"
	"strconv"
)

func CreatePost(c *gin.Context) {
	u := tool.GetUser(c)
	req := forumPB.CreatePostRequest{}
	tool.LoadMessage(c, &req)
	post := forum.Create(u.ID, req.Title, req.Content, req.Pictures)
	c.String(http.StatusOK, tool.DumpMessage(&forumPB.CreatePostReply{
		PostData: forum.Model2Data(post, u.ID),
	}))
}

func GetPost(c *gin.Context) {
	postIDParam := c.Query("post_id")
	postID, _ := strconv.Atoi(postIDParam)
	c.String(http.StatusOK, tool.DumpMessage(&forumPB.GetOnePostReply{
		Data: forum.GetOne(tool.GetUser(c).ID, uint(postID)),
	}))
}

func GetPostList(c *gin.Context) {
	posts := forum.GetAll(tool.GetUser(c).ID)
	c.String(http.StatusOK, tool.DumpMessage(&forumPB.GetPostListReply{
		Data: posts,
	}))
}

func LikePost(c *gin.Context) {
	req := forumPB.LikePostRequest{}
	tool.LoadMessage(c, &req)
	forum.Like(tool.GetUser(c).ID, uint(req.Id))
	c.String(http.StatusOK, tool.DumpMessage(&commonPB.EmptyMessage{}))
}

func UnLikePost(c *gin.Context) {
	req := forumPB.UnLikePostRequest{}
	tool.LoadMessage(c, &req)
	forum.UnLike(tool.GetUser(c).ID, uint(req.Id))
	c.String(http.StatusOK, tool.DumpMessage(&commonPB.EmptyMessage{}))
}
