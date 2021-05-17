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
