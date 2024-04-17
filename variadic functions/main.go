package main

import "fmt"

func main() {
	sumValue := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1)
	fmt.Println(sumValue)

}

func sum(args ...int) int {
	total := 0
	for _, n := range args {
		total += n
	}
	return total
}
