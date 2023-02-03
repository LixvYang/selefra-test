package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/payload", func(c *gin.Context) {
		buf := make([]byte, 1024*1024)
		n, _ := c.Request.Body.Read(buf)
		fmt.Println(string(buf[0:n]))
		resp := map[string]string{"hello": "world"}
		c.JSON(http.StatusOK, resp)
	})
	r.Run(":4567")
}
