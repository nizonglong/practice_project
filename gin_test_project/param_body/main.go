package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		bodyByts, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByts))
		firstName := c.Query("first_name")
		lastName := c.DefaultQuery("last_name", "last_default_name")
		c.String(http.StatusOK, string(bodyByts))
		c.String(http.StatusOK, "%s,%s", firstName, lastName)
	})
	r.Run() // listen and serve on 127.0.0.1:8080
}
