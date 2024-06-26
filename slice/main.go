package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 10, 20, 40, 80}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 11)

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[:2]), s[:2])
}
