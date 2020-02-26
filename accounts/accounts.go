package accounts

import (
	"errors"
	"fmt"
)

// account struct
type Account struct{
	owner string
	balance int
}

var errNoMoney = errors.New("Can't Withdraw")

// NewAccount create
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}


// a Account를 receiver라고 부른다
func (a *Account) Deposit(amount int){
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Widthdraw x amount from balance
func (a *Account) WithDraw(amount int) error {
	if a.balance < amount{
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// Changes the owner
func (a *Account) ChangeOwner(newOwner string){
	a.owner = newOwner
}

// Owner of Account
func (a Account) Owner()string{
	return a.owner
}

// struct를 print바로 해주는 method
// python의 __str__와 같다
func (a Account) String() string{
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}