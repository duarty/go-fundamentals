package main

import "fmt"

func main() {
	total := func() int {
		return sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1) * 2
	}()

	fmt.Println(total)

}

func sum(args ...int) int {
	total := 0
	for _, n := range args {
		total += n
	}
	return total
}
