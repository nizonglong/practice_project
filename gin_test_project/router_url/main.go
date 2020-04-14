package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": c.Param("name"),
			"id":   c.Param("id"),
		})
	})

	r.POST("/:name/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name_p": c.Param("name"),
			"id_p":   c.Param("id"),
		})
	})
	r.Run() // listen and serve on 127.0.0.1:8080
}
