package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
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
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test password=nizonglong sslmode=disable")

	if db == nil {
		fmt.Println("conn err=", err.Error())
	}
	db.LogMode(true)
	tid := c.Param("topic_id")
	topics := Topics{}
	db.Find(&topics, tid)
	c.JSON(http.StatusOK, topics)

	defer db.Close()
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
