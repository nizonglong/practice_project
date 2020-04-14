package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default_name")
		panic("test panic")
		c.String(http.StatusOK, "%s", name)
	})
	r.Run()
}
