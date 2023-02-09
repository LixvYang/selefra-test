package pr

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"selefra-demo/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

/**
 * @Author: LixvYang 2690688423@qq.com
 * @Date: 2023-02-08 14:13:06
 * @LastEditors: LixvYang 2690688423@qq.com
 * @LastEditTime: Do not edit
 * @FilePath: Do not edit
 * @Description:
 *
 * Copyright (c) 2023 by Lixvyang, All Rights Reserved.
 * @param {*gin.Context} c
 */
func MergePR(c *gin.Context) {
	var user model.User
	var githubPR GithubPR
	// x, _ := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(string(x))

	c.ShouldBind(&githubPR)
	// pr是否被合并
	// 检测对应的issue  #merged closed
	if githubPR.Action != "closed" && !githubPR.PullRequest.Merged {
		//不是合并直接退出
		return
	}

	// 查看issueNum
	issueIds, err := getIssueNum(&githubPR)
	if err != nil {
		c.JSON(http.StatusOK, "未绑定json")
		return
	}

	sumToken := sumIssueNums(issueIds)
	github_id := fmt.Sprint(githubPR.Sender.ID)

	user = model.User{
		GithubID:   github_id,
		GithubName: githubPR.Sender.Login,
		PublicKey:  "",
		AvatarUrl:  githubPR.Sender.AvatarURL,
		EmailLink:  "",
	}
	if user.CheckUser(&user); err != nil {
		err = user.CreateUser(&user)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}

	// sum 加总和
	err = user.IncrUserToken(&model.User{GithubID: github_id}, sumToken)
	if err != nil {
		log.Fatalf("增加失败")
		return
	}

	// 通过token num 总和 如果没有绑定
	// 合并成功
	// 查看一下参与人是否绑定钱包，绑定的话，则向pr发起者500token  给issure发起者50token
	// bind := user.CheckUserBind(&model.User{GithubID: github_id})
	// if !bind {
	// 没有绑定
	// 没有的话数据库暂时存储token等待github绑定钱包时一起绑定
	// user.IncrUserTempToken(&model.User{GithubID: github_id}, sumToken)
	// return
	// }

	// 绑定了的话就直接转账

}

func getIssueNum(githubPR *GithubPR) ([]string, error) {
	str := githubPR.PullRequest.Body
	reg := regexp.MustCompile(`#\d+`)
	issueIds := reg.FindAllString(str, -1)
	if len(issueIds) == 0 {
		return []string{}, errors.New("未找到 issueIds")
	}
	return issueIds, nil
}

// 根据issuenms #21 #22 总和所有的token nums
func sumIssueNums(issueIds []string) decimal.Decimal {
	var sum decimal.Decimal
	var issue model.Issue
	for i := 0; i < len(issueIds); i++ {
		i, err := issue.GetIssue(&model.Issue{IssueNumber: issueIds[i]})
		if err != nil {
			return decimal.Decimal{}
		}
		sum = sum.Add(i.TokenNum)
	}
	return sum
}
