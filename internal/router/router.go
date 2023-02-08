/*
 * @description:
 * @param:
 * @return:
 */
package router

import (
	"fmt"
	"selefra-demo/internal/api/issue"
	"selefra-demo/internal/api/pr"
	"selefra-demo/internal/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.POST("/api", gateway)

	_ = r.Run(utils.HttpPort)
}

func gateway(c *gin.Context) {
	fmt.Println("来了一个")
	githubEvent := c.Request.Header.Get("X-GitHub-Event")
	fmt.Println(githubEvent)

	switch {
	case githubEvent == "issues":
		issue.AddIssue(c)
	case githubEvent == "pull_request":
		pr.MergePR(c)
	}
}
