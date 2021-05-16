package qa

import "ncutbbs/model"

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
