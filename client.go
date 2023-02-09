package main

import (
	"fmt"
	"log"
	"selefra-demo/internal/token"
)

func main() {
	token.Init()
	t := token.T{}
	// fmt.Println(t.Name())
	// fmt.Println(t.Symbol())
	// fmt.Println(t.BalanceOf("0x0eecD003Aa72554354527BeD08dC88f971d881DF"))
	// fmt.Println(t.Decimals())
	// fmt.Println(t.Owner())
	// err := t.T("0xEf42420F5d2815CbB2700d03D527F0F89bdA9503", 100000000)
	err := t.T("nihao", 1000)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(11)
}
