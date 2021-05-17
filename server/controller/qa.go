package controller

import (
	"github.com/gin-gonic/gin"
	"ncutbbs/module/qa"
	commonPB "ncutbbs/proto/common"
	qaPB "ncutbbs/proto/qa"
	"ncutbbs/server/controller/tool"
	"net/http"
	"strconv"
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

func GetAnswer(c *gin.Context) {
	questionID, _ := strconv.Atoi(c.Query("question_id"))
	c.String(http.StatusOK, tool.DumpMessage(&qaPB.GetAnswerReply{
		Data: qa.GetAnswers(uint(questionID), tool.GetUser(c).ID),
	}))
}

func CreateAnswer(c *gin.Context) {
	req := qaPB.CreateAnswerRequest{}
	tool.LoadMessage(c, &req)
	qa.CreateAnswer(uint(req.QuestionId), tool.GetUser(c).ID, req.Content)
	c.String(http.StatusOK, tool.DumpMessage(&commonPB.EmptyMessage{}))
}

func LikeAnswer(c *gin.Context) {
	req := qaPB.LikeAnswerRequest{}
	tool.LoadMessage(c, &req)
	qa.LikeAnswer(uint(req.Id), tool.GetUser(c).ID)
	c.String(http.StatusOK, tool.DumpMessage(&commonPB.EmptyMessage{}))
}

func UnlikeAnswer(c *gin.Context) {
	req := qaPB.UnLikeAnswerRequest{}
	tool.LoadMessage(c, &req)
	qa.UnlikeAnswer(uint(req.Id), tool.GetUser(c).ID)
	c.String(http.StatusOK, tool.DumpMessage(&commonPB.EmptyMessage{}))
}
