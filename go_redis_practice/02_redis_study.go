package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 打印redis所有k-v
func printValues(keys [][]byte) {
	for _, v := range keys { // 忽略 index
		val, _ := redis.String(c.Do("GET", v))
		fmt.Printf("key %s : %s\n", v, val)
	}
}

var c, err = redis.Dial("tcp", "127.0.0.1:6379")

func main() {

	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	// 获取所有的key
	keys, err := redis.ByteSlices(c.Do("keys", "*"))
	if err != nil {
		fmt.Println("redis get keys failed:", err)
	} else {
		printValues(keys)
	}

}
