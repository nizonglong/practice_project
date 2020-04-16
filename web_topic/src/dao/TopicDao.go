package dao

import (
	"github.com/gin-gonic/gin"
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
	//c.String(http.StatusOK, "获取topic_id=%s的帖子", c.Param("topic_id"))
	c.JSON(http.StatusOK, CreateTopic(1, "title 1"))

}

func GetTopicList(c *gin.Context) {
	//if "" == c.Query("username") {
	//	c.String(http.StatusOK, "获取帖子列表")
	//} else {
	//	c.String(http.StatusOK, "获取用户%s的帖子列表", c.Query("username"))
	//}

	query := TopicQuery{}
	err := c.BindQuery(&query)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误：%s", err.Error())
	} else {
		c.JSON(http.StatusOK, query)
	}
}

func NewTopic(c *gin.Context) {
	topic := Topic{}
	err := c.BindJSON(&topic)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误：%s", err.Error())
	} else {
		c.JSON(http.StatusOK, topic)
	}
}

func DeleteTopic(c *gin.Context) {
	c.String(http.StatusOK, "删除帖子")
}
