package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "hello test")
	})

	autotls.Run(r, "www.itpp.tk")
}
