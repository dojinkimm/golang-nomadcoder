package main

import (
	"fmt"
	"github.com/dojinkimm/golang_study/accounts"
)


func main() {
	account := accounts.NewAccount("Henry")
	account.Deposit(10)
	fmt.Println(account)
}
