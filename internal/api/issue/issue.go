package issue

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"selefra-demo/internal/model"
	"selefra-demo/internal/utils"
	"selefra-demo/internal/utils/email"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func AddIssue(c *gin.Context) {
	// var user model.User
	var issue model.Issue
	var githubIssue GithubIssue
	c.ShouldBind(&githubIssue)
	// fmt.Println("来这里")
	// x, _ := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(string(x))
	if githubIssue.Action == "labeled" {
		return
	}

	// read token
	github_id, token_num, err := getTokenNums(&githubIssue)
	if err != nil {
		log.Fatalf("getTokenNums err !")
		c.JSON(http.StatusOK, "error: get github_id or token error!"+err.Error())
		return
	}

	// 检测issue发起者的token是否足够，不足则关闭issue
	// 给我传递过来github_id
	// if !checkTokenEnough(github_id, token_num) {
	// 	c.JSON(http.StatusOK, "error: token not enough.")
	// 	// close issue
	// 	return
	// }

	// 充足则从issue发起者扣除500token，数据库记录一下需要扣除的金额      ，等到issue解决或关闭时，再划转
	// if err = user.DescUserTokenNum(&model.User{GithubID: github_id}, token_num); err != nil {
	// 	return
	// }

	// 记录对应的Issure
	issue.Body = githubIssue.Issue.Body
	issue.TokenNum = token_num
	// 这里是demo
	// TODO
	issue.IssueNumber = "#" + fmt.Sprint(githubIssue.Issue.Number)
	issue.Uid = github_id
	if err = issue.CreateIssue(&issue); err != nil {
		// 创建Issue错误
		log.Fatalf("CreateIssue err !", err)
		return
	}

	// 如果用户不存在,则创建用户
	user := model.User{
		GithubID:   github_id,
		GithubName: githubIssue.Sender.Login,
		PublicKey:  "",
		AvatarUrl:  githubIssue.Sender.AvatarURL,
		EmailLink:  "",
	}
	err = user.CheckUser(&user)
	if err != nil {
		user.CreateUser(&user)
	}

	// err = user.CreateUser(&user); if err != nil {

	// }
	// send this issure to every one who participant
	// sendParticipantEmail(githubIssue)

	c.JSON(http.StatusOK, "success!")
	return
}

func getTokenNums(githubIssue *GithubIssue) (github_id string, token decimal.Decimal, err error) {
	str := githubIssue.Issue.Body
	_, after, _ := strings.Cut(str, "token 数目")
	buf, _, _ := strings.Cut(after, "### Golang 版本")
	fmt.Println(strings.ReplaceAll(buf, "\n", ""))
	token, _ = decimal.NewFromString(strings.ReplaceAll(buf, "\n", ""))
	// tokenNum := strings.Split(str, "\n\n### token 数目\n\n")[1]
	// tokenNum = strings.Split(tokenNum, "\n\n")[0]
	// token, err = decimal.NewFromString(tokenNum)
	// if err != nil {
	// 	return "", decimal.Decimal{}, errors.New("Error" + err.Error())
	// }

	return fmt.Sprint(githubIssue.Issue.User.ID), token, nil
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
