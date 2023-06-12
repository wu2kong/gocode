package main

import (
	"fmt"
	load_balance "gocode/pkg/load_balance"
)

func main() {
	ldb := load_balance.LoadBalanceFactory(load_balance.LbRandom)
	ldb.Add("ac", "aa")
	fmt.Println(ldb.Get("ac"))
	fmt.Printf("aa")
}
