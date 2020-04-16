package main

import (
	"github.com/gin-gonic/gin"
	. "practice_project/web_topic/src/dao"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1/topics")
	// v1的请求放在代码块里，更加容易理解和阅读
	{
		v1.GET("", GetTopicList)

		v1.GET("/:topic_id", GetTopicDetail)

		v1.Use(MustLogin())
		{
			v1.POST("", NewTopic)

			v1.DELETE("/:topic_id", DeleteTopic)
		}

	}

	router.Run(":8080")
}
