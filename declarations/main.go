package main

import "fmt"

const a = "hello world"

var b bool
var c bool = true

var (
	d bool
	e int
	f string
	g float64 = 0.01
)

func main() {
	d = true
	h := c == d
	fmt.Println(b)
	fmt.Println(h)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	f = "set value"
	fmt.Println(f)
	fmt.Println(g)

}
