package dao

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	. "practice_project/web_topic/src/cache"
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
	c.Set("dbResult", topics)
	//c.JSON(http.StatusOK, topics)

	//conn := RedisDefaultPool.Get()
	//redisKey := "topic_" + tid
	//defer conn.Close()
	//ret, err := redis.Bytes(conn.Do("get", redisKey))
	//if err != nil { // 缓存里没有
	//    DBHelper.Find(&topics, tid)
	//    retData, _ := ffjson.Marshal(topics)
	//    if topics.TopicID == 0 { // 数据库没有匹配到
	//        conn.Do("setex", redisKey, 20, retData)
	//    } else { // 正常数据50s缓存
	//        conn.Do("setex", redisKey, 50, retData)
	//    }
	//
	//    c.JSON(http.StatusOK, topics)
	//    log.Println("从数据库读取")
	//} else { // 缓存有值
	//    ffjson.Unmarshal(ret, &topics)
	//    c.JSON(http.StatusOK, topics)
	//    log.Println("从缓存读取")
	//}
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
