package qa

import (
	"ncutbbs/model"
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
}

func LikeAnswer(answerID, userID uint) {
	al := model.AnswerLike{
		UserID:   userID,
		AnswerID: answerID,
	}
	model.DB.Create(&al)
}

func UnlikeAnswer(answerID, userID uint) {
	model.DB.Where("answer_id = ? AND user_id = ?", answerID, userID).Delete(&model.AnswerLike{})
}
