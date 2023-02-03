package main

import (
	"fmt"
	"strings"
)

func main() {
	str := `### .... 版本\n\n11\n\n### token 数目\n\n500\n\n### Golang 版本\n\ngo 0.20.0\n\n### 是否依旧存在\n\n可以\n\n### bug描述\n\n1\n\n### 修改建议\n\n_No response_`
	fmt.Println(str)
	_, after, _ := strings.Cut(str, "token 数目")
	buf, _, _ := strings.Cut(after, "### Golang 版本")
	// fmt.Println(strings.Trim(buf, ""))
	fmt.Println(strings.TrimSpace(buf))
	// fmt.Println(aft)
	// fmt.Println(buf)
	// fmt.Println(after[4:])
	return
}
