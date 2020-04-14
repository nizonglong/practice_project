package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", testing)
	r.POST("/testing", testing)
	r.Run(":8080") // listen and serve on 127.0.0.1:8080
}

func testing(c *gin.Context) {
	var person Person
	// 这里是根据请求的content type来做不同binding操作
	if err := c.ShouldBind(&person); err == nil {
		c.String(http.StatusOK, "%v", person)
	} else {
		c.String(http.StatusInternalServerError, "person bind error : %v", err)
	}
}
