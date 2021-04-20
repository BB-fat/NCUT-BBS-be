package controller

import (
	"github.com/gin-gonic/gin"
	"ncutbbs/module/forum"
	forumPB "ncutbbs/proto/forum"
	"ncutbbs/server/controller/tool"
	"net/http"
)

func CreatePost(c *gin.Context) {
	u := tool.GetUser(c)
	req := forumPB.CreatePostRequest{}
	tool.LoadMessage(c, &req)
	post := forum.Create(u.ID, req.Title, req.Content, req.Pictures)
	c.String(http.StatusOK, tool.DumpMessage(&forumPB.CreatePostReply{
		PostData: forum.Model2Data(post),
	}))
}

func GetPostList(c *gin.Context) {
	posts := forum.GetAll()
	c.String(http.StatusOK, tool.DumpMessage(&forumPB.GetPostListReply{
		Data: posts,
	}))
}
