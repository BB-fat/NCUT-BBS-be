package controller

import (
	"github.com/gin-gonic/gin"
	"ncutbbs/module/news"
	newsPB "ncutbbs/proto/news"
	"ncutbbs/server/controller/tool"
	"net/http"
)

func GetNews(c *gin.Context) {
	c.String(http.StatusOK, tool.DumpMessage(&newsPB.GetNewsReply{
		Data: news.GetNewsList(tool.GetUser(c).ID),
	}))
}
