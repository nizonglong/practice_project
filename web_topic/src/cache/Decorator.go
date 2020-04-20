package cache

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/pquerna/ffjson/ffjson"
	"log"
	"net/http"
)

func CacheDecorator(h gin.HandlerFunc, param string, redKayPattern string, empty interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		// redis判断
		getID := context.Param(param)
		redisKey := fmt.Sprintf(redKayPattern, getID)
		conn := RedisDefaultPool.Get()
		defer conn.Close()
		ret, err := redis.Bytes(conn.Do("get", redisKey))
		if err != nil { // 缓存里没有
			h(context) // 执行目标方法
			dbResult, exists := context.Get("dbResult")
			if !exists {
				dbResult = empty
			}

			retData, _ := ffjson.Marshal(dbResult)
			conn.Do("setex", redisKey, 20, retData)
			context.JSON(http.StatusOK, dbResult)
			log.Println("从数据库读取")
		} else { // 缓存有，直接抛出
			log.Println("从缓存读取")
			ffjson.Unmarshal(ret, &empty)
			context.JSON(http.StatusOK, empty)
		}
	}
}
