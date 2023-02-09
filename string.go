package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	str := "Fix bugs...\r\n\r\n#23 #2 "
// 	// fmt.Println(str)
// 	// _, after, _ := strings.Cut(str, "token 数目")
// 	// buf, _, _ := strings.Cut(after, "### Golang 版本")
// 	// fmt.Println(strings.TrimSpace(buf))
// 	// fmt.Println(strings.ReplaceAll(buf, "\n", ""))
// 	// fmt.Println(aft)
// 	// fmt.Println(buf)
// 	// fmt.Println(after[4:])

// 	reg := regexp.MustCompile(`#\d+`)
// 	fmt.Println(reg.FindAllString(str, -1))
// 	return
// }

func main() {
	str := "### 版本\n\n11\n\n### token 数目\n\n500\n\n### Golang 版本\n\ngo 0.20.0\n\n### 是否依旧存在\n\n可以\n\n### bug描述\n\n1\n\n### 修改建议\n\n_No response_"
	// tokenNum := strings.Split(str, "\n\n### token 数目\n\n")[1]
	// tokenNum = strings.Split(tokenNum, "\n\n")[0]
	// tokenNumb, err := strconv.Atoi(tokenNum)
	// if err != nil {
	// 	log.Fatalf("error: ", err)
	// }
	// fmt.Println(tokenNumb)

	_, after, found := strings.Cut(str, "token 数目")
	if !found {
		return
	}
	buf, _, _ := strings.Cut(after, "### Golang 版本")
	fmt.Println(strings.ReplaceAll(buf, "\n", ""))
}
