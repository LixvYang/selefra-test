package main

import (
	"fmt"
	"io/ioutil"
	"selefra-demo/issue"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/payload", func(c *gin.Context) {
		x, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Printf("%s", string(x))
		var githubIssue issue.GithubIssue
		c.ShouldBind(&githubIssue)
		fmt.Println(githubIssue)
		str := githubIssue.Issue.Body
		_, after, found := strings.Cut(str, "token 数目")
		if !found {
			c.JSON(200, "error: Not found tokens")
			return
		}
		buf, _, _ := strings.Cut(after, "### Golang 版本")
		fmt.Println(strings.ReplaceAll(buf, "\n", ""))
		return
	})
	r.Run(":4567")
}
