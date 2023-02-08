package issue

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"selefra-demo/internal/model"
	"selefra-demo/internal/utils"
	"selefra-demo/internal/utils/email"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func AddIssue(c *gin.Context) {
	var user model.User
	var githubIssue GithubIssue
	c.ShouldBind(&githubIssue)

	// read token
	github_id, token_num, err := getTokenNums(githubIssue)
	if err != nil {
		c.JSON(http.StatusOK, "error: get github_id or token error!"+err.Error())
		return
	}

	// 检测issue发起者的token是否足够，不足则关闭issue
	// 给我传递过来github_id
	if !checkTokenEnough(github_id, token_num) {
		c.JSON(http.StatusOK, "error: token not enough.")
		// close issue
		return
	}

	// 充足则从issue发起者扣除500token，数据库记录一下需要扣除的金额      ，等到issue解决或关闭时，再划转
	if err = user.AddUserPreDeductNum(&model.User{GithubID: github_id}, token_num); err != nil {
		return
	}

	// send this issure to every one who participant
	sendParticipantEmail(githubIssue)

	c.JSON(http.StatusOK, "success!")
	return
}

func getTokenNums(githubIssue GithubIssue) (github_id string, token decimal.Decimal, err error) {
	str := githubIssue.Issue.Body
	_, after, found := strings.Cut(str, "token 数目")
	if !found {
		return "", decimal.Decimal{}, errors.New("Not found tokens")
	}
	buf, _, _ := strings.Cut(after, "### Golang 版本")
	fmt.Println(strings.ReplaceAll(buf, "\n", ""))

	token, err = decimal.NewFromString(buf)
	if err != nil {
		return "", decimal.Decimal{}, errors.New("Error" + err.Error())
	}

	return string(githubIssue.Issue.User.ID), token, nil
}

// 检测issue发起者token是否足够
func checkTokenEnough(github_id string, token decimal.Decimal) bool {
	var u model.User
	userToken, err := u.GetUserTokenNum(&model.User{GithubID: github_id})
	if err != nil {
		return false
	}

	if userToken.LessThan(token) {
		return false
	}

	return true
}

// 给参与者发送邮件
func sendParticipantEmail(githubIssue GithubIssue) {
	var user model.User
	users, _ := user.ListUsers()
	MailTo := email.ConvertToUsers(users)

	options := &email.Options{
		MailHost: utils.MailHost,
		MailPort: utils.MailPort,
		MailUser: utils.MailUser,
		MailPass: utils.MailPass,
		MailTo:   MailTo,
		Subject:  "subject",
		Body:     GetEmailBody(githubIssue),
	}
	if err := email.Send(options); err != nil {
		fmt.Println("send email error: ", err)
	}
}

// 先定义死
func GetEmailBody(githubIssue GithubIssue) string {
	return "body"
}
