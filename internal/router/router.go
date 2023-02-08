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
	a := r.Group("/api")

	{
		a.POST("/issue/add", issue.AddIssue)
		a.POST("/pr/add", pr.AddPR)
		a.POST("/pr/merge", pr.MergePR)
	}

	_ = r.Run(utils.HttpPort)
}
