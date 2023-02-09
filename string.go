package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "Fix bugs...\r\n\r\n#23 #2 "
	// fmt.Println(str)
	// _, after, _ := strings.Cut(str, "token 数目")
	// buf, _, _ := strings.Cut(after, "### Golang 版本")
	// fmt.Println(strings.TrimSpace(buf))
	// fmt.Println(strings.ReplaceAll(buf, "\n", ""))
	// fmt.Println(aft)
	// fmt.Println(buf)
	// fmt.Println(after[4:])

	reg := regexp.MustCompile(`#\d+`)
	fmt.Println(reg.FindAllString(str, -1))
	return
}
