package main

import (
	"fmt"
	"log"
	"selefra-demo/internal/token"
)

func main() {
	token.Init()
	t := token.T{}
	nums := t.BalanceOf("0xef42420f5d2815cbb2700d03d527f0f89bda9503")
	nums = t.BalanceOf("0x0eecD003Aa72554354527BeD08dC88f971d881DF")
	fmt.Println(nums)
	err := t.T("0xef42420f5d2815cbb2700d03d527f0f89bda9503", 1)
	if err != nil {
		log.Fatalf(err.Error())
	}

	nums = t.BalanceOf("0xef42420f5d2815cbb2700d03d527f0f89bda9503")
	fmt.Println(nums)
}
