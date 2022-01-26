package accounts

import (
	"errors"
	"fmt"
)

// use uppercase to export anything (event fields in struct !)
type Account struct {
	owner   string
	balance int
}

// error variable's name should be err- prefixed
var errNoMoney = errors.New("not enough balance in your account")

// normal constructor pattern in Go
// returing the real address of the object
// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // returns address of account, not a copy
}

// (a *Account) is receiver
// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
// just use copy of Account, not *Account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}

	a.balance -= amount
	return nil
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// Overrides Go's built-in String() method
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account \nHas: ", a.Balance())
}
