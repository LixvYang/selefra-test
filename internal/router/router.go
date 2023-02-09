/*
 * @description:
 * @param:
 * @return:
 */
/*
 * @description:
 * @param:
 * @return:
 */
package router

import (
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
	githubEvent := c.Request.Header.Get("X-GitHub-Event")
	switch githubEvent {
	case "issues":
		issue.AddIssue(c)
	case "pull_request":
		pr.MergePR(c)
	}
}
