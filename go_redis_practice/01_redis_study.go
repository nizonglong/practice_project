package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "test", "go-redis")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	val, err := redis.String(c.Do("GET", "test"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get test: %v \n", val)
	}
}
