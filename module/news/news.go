package news

import (
	"google.golang.org/protobuf/proto"
	"ncutbbs/model"
	newsPB "ncutbbs/proto/news"
	"ncutbbs/server/controller/tool"
)

func CreateNews(newsType newsPB.NewsType, title, content string, userID uint, data proto.Message) {
	n := model.News{
		Type:     int(newsType),
		Title:    title,
		Content:  content,
		UserID:   userID,
		Data:     tool.DumpMessage(data),
		Received: 0,
	}
	model.DB.Create(&n)
}

func GetNewsList(userID uint) []*newsPB.NewsData {
	var list []model.News
	model.DB.Where("user_id = ? AND received = 0", userID).Find(&list)
	data := make([]*newsPB.NewsData, len(list), len(list))
	for i := 0; i < len(list); i++ {
		data[i] = list[i].ToData()
	}
	//TODO receive 设置为1
	return data
}
