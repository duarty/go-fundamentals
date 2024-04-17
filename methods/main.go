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

func (u *User) Disable() {
	u.Active = false
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

	user.Address.City = "Dublin"
	fmt.Println(user)
}
