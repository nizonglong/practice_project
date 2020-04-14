package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/user/*action", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	r.Run() // listen and serve on 127.0.0.1:8080
}
