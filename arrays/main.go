package main

import "fmt"

func main() {
	var arr [4]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 100

	fmt.Println(len(arr))
	fmt.Println(arr[len(arr)-1])

	for i, v := range arr {
		fmt.Printf("O valor do índice %d é %d\n", i, v)
	}
}
