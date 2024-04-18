package main

func main() {

	arr := []int{1, 2, 3, 4, 5}

	for i := 0; i < 10; i++ {
		println(i)
	}

	for index, value := range arr {
		println(index, value)
	}

	for _, value := range arr {
		println(value)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	for {
		println("...")
	}
}
