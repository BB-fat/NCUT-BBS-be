package controller

import (
	"github.com/gin-gonic/gin"
	"ncutbbs/module/qa"
	commonPB "ncutbbs/proto/common"
	qaPB "ncutbbs/proto/qa"
	"ncutbbs/server/controller/tool"
	"net/http"
)

func CreateQuestion(c *gin.Context) {
	u := tool.GetUser(c)
	req := qaPB.CreateQuestionRequest{}
	tool.LoadMessage(c, &req)
	q := qa.CreateQuestion(u.ID, req.Title, req.Content, req.Pictures)
	c.String(http.StatusOK, tool.DumpMessage(&qaPB.CreateQuestionReply{
		Data: q.ToData(),
	}))
}

func GetQuestionList(c *gin.Context) {
	c.String(http.StatusOK, tool.DumpMessage(&qaPB.GetQuestionListReply{
		Data: qa.GetAllQuestion(),
	}))
}

func AddQuestionViews(c *gin.Context) {
	req := qaPB.AddQuestionViewsRequest{}
	tool.LoadMessage(c, &req)
	qa.AddViews(uint(req.Id))
	c.String(http.StatusOK, tool.DumpMessage(&commonPB.EmptyMessage{}))
}
