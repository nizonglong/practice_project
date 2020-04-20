package dao

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	. "practice_project/web_topic/src"
	. "practice_project/web_topic/src/model"
)

func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, status := c.GetQuery("token"); !status {
			c.String(http.StatusUnauthorized, "缺少token参数")
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func GetTopicDetail(c *gin.Context) {
	tid := c.Param("topic_id")
	topics := Topics{}
	DBHelper.Find(&topics, tid)
	//DBHelper.Find(&topics, "topic_id=1")
	c.JSON(http.StatusOK, topics)
}

func GetTopicList(c *gin.Context) {
	query := TopicQuery{}
	err := c.BindQuery(&query)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误：%s", err.Error())
	} else {
		c.JSON(http.StatusOK, query)
	}
}

// 新增单条帖子
func NewTopic(c *gin.Context) {
	topic := Topics{}
	err := c.BindJSON(&topic)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误：%s", err.Error())
	} else {
		c.JSON(http.StatusOK, topic)
	}
}

// 新增多条帖子
func NewTopics(c *gin.Context) {
	topics := Topics{}
	err := c.BindJSON(&topics)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误：%s", err.Error())
	} else {
		c.JSON(http.StatusOK, topics)
	}
}

func DeleteTopic(c *gin.Context) {
	c.String(http.StatusOK, "删除帖子")
}
