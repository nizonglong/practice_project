package main

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/go-playground/validator.v8"
	"log"
	"net/http"
	. "practice_project/web_topic/src"
	. "practice_project/web_topic/src/cache"
	. "practice_project/web_topic/src/dao"
	. "practice_project/web_topic/src/model"
	"time"
)

func main2() {
	conn := RedisDefaultPool.Get()
	ret, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(ret)
}

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

		v1.GET("/:topic_id", CacheDecorator(GetTopicDetail, "topic_id", "topic_%s", TopicClass{}))

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

	//router.Run()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go (func() { // 启动web服务
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("服务器启动失败: ", err)
		}
	})()

	go (func() {
		InitDB()
	})()

	ServerNotify()
	// 这里还可以做一些 释放或者善后工作，此处略去
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal("服务器关闭")
	}
	log.Println("服务器优雅的退出")
}
