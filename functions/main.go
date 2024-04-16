package main

import (
	"errors"
	"fmt"
)

func main() {
	x := 1
	y := 2
	fmt.Println(sum(x, y))
	z := sum(x, y) * 3
	fmt.Println("z value is", z)
	v := &z
	div, err := divideBy(v, 3)
	if err != nil {
		fmt.Println(div, err)
		return
	}
	fmt.Printf("%d divided by %d is %d\n", *v, z, div)
	fmt.Println("new v and z value is", *v, z)

}

func sum(a, b int) int {
	return a + b
}

func divideBy(a *int, b int) (int, error) {
	*a = 3
	if *a%b > 0 {
		return 0, errors.New("Resto da divisão maior que zero")
	}
	return *a / b, nil
}
