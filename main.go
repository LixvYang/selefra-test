package main

import (
	"fmt"
	"selefra-demo/issue"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/payload", func(c *gin.Context) {
		var githubIssue issue.GithubIssue
		c.ShouldBind(&githubIssue)
		// strings.
		// body, err := ioutil.ReadAll(c.Request.Body)
		// if err != nil {
		// 	// Handle error
		// 	log.Fatal(err)
		// }
		fmt.Println(githubIssue.Issue.Body)
	})
	r.Run(":4567")
}
