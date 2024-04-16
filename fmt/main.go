package main

import "fmt"

var (
	d bool
	z float64 = 0.16666
)

func main() {
	d = true
	fmt.Printf("d type is %T", d)
	fmt.Printf("z value is %E", z)

}
