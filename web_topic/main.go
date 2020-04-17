package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/go-playground/validator.v8"
	. "practice_project/web_topic/src"
	. "practice_project/web_topic/src/dao"
)

func main() {
	router := gin.Default()
	// 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("topicurl", TopicUrl)
		// 验证长度
		_ = v.RegisterValidation("topics", TopicsValidate)
	}

	/**
	 * 单条帖子的处理
	 */
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

	/**
	 * 多条帖子处理
	 */
	v2 := router.Group("/v1/mtopics")
	// v1的请求放在代码块里，更加容易理解和阅读
	{
		v2.Use(MustLogin())
		{
			v2.POST("", NewTopics)
		}

	}

	router.Run(":8080")

}
