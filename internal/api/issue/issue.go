package issue

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddIssue(c *gin.Context) {
	// read token
	x, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("%s", string(x))
	var githubIssue GithubIssue
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

	// 检测issue发起者的token是否足够，不足则关闭issue

	// 充足则从issue发起者扣除500token，数据库记录一下需要扣除的金额，等到issue解决或关闭时，再划转  

	// send this issure to every one who participant


	return
}
