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

func main() {
	user := User{
		Name:   "José Duarte",
		Age:    33,
		Active: false,
		Address: Address{
			Street:     "RLBD",
			Number:     599,
			City:       "BVB",
			State:      "RR",
			PostalCode: "69300000",
		},
	}

	user.Active = false
	user.Address.City = "Dublin"
	fmt.Println(user)
}
