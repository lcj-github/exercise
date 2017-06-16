package main

import "fmt"
import "time"
import "math/rand"

//http://blog.csdn.net/yugongpeng_blog/article/details/43530771
// 目前只是 没处理 并发问题
type Account interface {
	Withdraw(uint)
	Deposit(uint)
	Balance() int
}
type Bank struct {
	account Account
}

func NewBank(account Account) *Bank {
	return &Bank{account: account}
}
func (bank *Bank) Withdraw(amount uint, actor_name string) {
	fmt.Println("[-]", amount, actor_name)
	bank.account.Withdraw(amount)
}
func (bank *Bank) Deposit(amount uint, actor_name string) {
	fmt.Println("[+]", amount, actor_name)
	bank.account.Deposit(amount)
}
func (bank *Bank) Balance() int {
	return bank.account.Balance()
}

type SimpleAccount struct {
	balance int
}

func NewSimpleAccount(balance int) *SimpleAccount {
	return &SimpleAccount{balance: balance}
}
func (acc *SimpleAccount) Deposit(amount uint) {
	acc.setBalance(acc.balance + int(amount))
}
func (acc *SimpleAccount) Withdraw(amount uint) {
	if acc.balance >= int(amount) {
		acc.setBalance(acc.balance - int(amount))
	} else {
		panic("杰克穷死")
	}
}
func (acc *SimpleAccount) Balance() int {
	return acc.balance
}
func (acc *SimpleAccount) setBalance(balance int) {
	acc.add_some_latency() //增加一个延时函数，方便演示
	acc.balance = balance
}
func (acc *SimpleAccount) add_some_latency() {
	<-time.After(time.Duration(rand.Intn(100)) * time.Millisecond)
}
func main() {
	balance := 80
	b := NewBank(NewSimpleAccount(balance))

	fmt.Println("初始化余额", b.Balance())

	b.Withdraw(30, "马伊琍")

	fmt.Println("-----------------")
	fmt.Println("剩余余额", b.Balance())
}
