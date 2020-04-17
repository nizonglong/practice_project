package model

import (
	"time"
)

type Topics struct {
	TopicID         int       `json:"id"`
	TopicTitle      string    `json:"title" binding:"min=4,max=20"`
	TopicShortTitle string    `json:"stitle" binding:"required"`
	UserIP          string    `json:"ip" binding:"ipv4"`
	TopicScore      int       `json:"score" binding:"gt=5"`
	TopicUrl        string    `json:"url" binding:"omitempty,topicurl"`
	TopicDate       time.Time `json:"date" binding:"required"`
}

type TopicArray struct {
	TopicList     []Topics `json:"topics" binding:"gt=0,lt=3,topics,dive"`
	TopicListSize int      `json:"size"`
}

type TopicClass struct {
	ClassId     int
	ClassName   string
	ClassRemark string
}

func CreateTopic(id int, title string) Topics {
	return Topics{TopicID: id, TopicTitle: title}
}

type TopicQuery struct {
	UserName string `json:"username" form:"username"`
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"pagesize" form:"pagesize"`
}

/**
多条Topic测试json
{
  "topics":
  [
      {"title":"abcd","stitle":"abc","ip":"127.0.0.1","score":6,"url":"aaa"},
      {"title":"abcd2","stitle":"abc2","ip":"127.0.0.1","score":8,"url":"aaa2"}
  ],
  "size": 2
}
*/
