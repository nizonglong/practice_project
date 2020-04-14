package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{
			"127.0.0.2",
		}

		flag := false
		clientIP := c.ClientIP()
		for _, host := range ipList {
			if clientIP == host {
				flag = true
				break
			}
		}

		if !flag {
			c.String(http.StatusUnauthorized, "%s, not in iplist", clientIP)
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(IPAuthMiddleware())
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "hello test")
	})
	r.Run()
}
