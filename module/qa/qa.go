package qa

import (
	"fmt"
	"ncutbbs/model"
	"ncutbbs/module/news"
	newsPB "ncutbbs/proto/news"
	qaPB "ncutbbs/proto/qa"
)

func CreateQuestion(authorID uint, title, content string, pictures string) model.Question {
	question := model.Question{
		AuthorID: authorID,
		Title:    title,
		Views:    1,
		Content:  content,
		Pictures: pictures,
	}
	model.DB.Create(&question)
	return question
}

func DeleteQuestion(ID uint) bool {
	question := model.Post{
		ID: ID,
	}
	res := model.DB.Delete(&question)
	if res.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

func GetAllQuestion() []*qaPB.QuestionData {
	var questions []model.Question
	model.DB.Find(&questions)
	var data []*qaPB.QuestionData
	for i := 0; i < len(questions); i++ {
		data = append(data, questions[i].ToData())
	}
	return data
}

func AddViews(questionID uint) {
	var question model.Question
	model.DB.First(&question, questionID)
	question.Views++
	model.DB.Save(&question)
}

func GetAnswers(questionID, userID uint) []*qaPB.AnswerData {
	var data []model.Answer
	model.DB.Where("question_id = ?", questionID).Find(&data)
	list := make([]*qaPB.AnswerData, len(data), len(data))
	for i := 0; i < len(data); i++ {
		list[i] = data[i].ToData(userID)
	}
	return list
}

func CreateAnswer(questionID, userID uint, content string) {
	answer := model.Answer{
		AuthorID:   userID,
		QuestionID: questionID,
		Content:    content,
	}
	model.DB.Create(&answer)

	q := model.Question{ID: questionID}
	model.DB.First(&q)
	if q.AuthorID != userID {
		u := model.User{ID: userID}
		model.DB.First(&u)
		news.CreateNews(newsPB.NewsType_QA, fmt.Sprintf("%s回答了你提出的问题", u.AccountName), content, q.AuthorID, q.ToData())
	}
}

func LikeAnswer(answerID, userID uint) {
	al := model.AnswerLike{
		UserID:   userID,
		AnswerID: answerID,
	}
	model.DB.Create(&al)

	a := model.Answer{ID: answerID}
	model.DB.First(&a)
	q := model.Question{ID: a.QuestionID}
	model.DB.First(&q)
	if a.AuthorID != userID {
		u := model.User{ID: userID}
		model.DB.First(&u)
		news.CreateNews(newsPB.NewsType_QA, fmt.Sprintf("%s给你的回答点赞", u.AccountName), "轻点查看详情", a.AuthorID, q.ToData())
	}
}

func UnlikeAnswer(answerID, userID uint) {
	model.DB.Where("answer_id = ? AND user_id = ?", answerID, userID).Delete(&model.AnswerLike{})
}
