package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Active bool
	Address
}

type Address struct {
	Street     string
	Number     int
	City       string
	State      string
	PostalCode string
}

type Person interface {
	Disable()
	Enable()
}

func (u *User) Disable() {
	u.Active = false
}

func (u *User) Enable() {
	u.Active = true
}

func Diactivation(person Person) {
	person.Disable()
}

func Activation(person Person) {
	person.Enable()
}

func main() {
	user := User{
		Name:   "José Duarte",
		Age:    33,
		Active: true,
		Address: Address{
			Street:     "RLBD",
			Number:     599,
			City:       "BVB",
			State:      "RR",
			PostalCode: "69300000",
		},
	}
	fmt.Println(user)
	user.Disable()
	fmt.Println(user)
	Activation(&user)
	user.Address.City = "Dublin"
	fmt.Println(user)
}
