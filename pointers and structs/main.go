package main

import "fmt"

type User struct {
	name string
}

type Acc struct {
	balance float64
}

func (a Acc) simulation(v float64) float64 {
	a.balance += v
	return a.balance
}

func (a *Acc) borrow(v float64) float64 {
	a.balance += v
	return a.balance
}

func (u *User) walk() {
	u.name = "WW"
	fmt.Printf("%v is  \n", u.name)
}

func NewAcc() *Acc {
	return &Acc{balance: 0.0}
}

func main() {
	user := User{
		name: "Duarte",
	}
	user.walk()
	fmt.Printf("client name is: %v\n", user.name)

	acc := Acc{
		balance: 1.234,
	}
	acc.simulation(1.000)
	println(acc.balance)
	acc.borrow(1.000)
	println(acc.balance)

	newAcc := NewAcc()
	println(newAcc.balance)
	newAcc.balance = 1.000
	println(newAcc.balance)

}
