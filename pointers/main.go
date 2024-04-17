package main

func main() {
	a := 10
	println(a)
	var pointerA *int = &a
	println(pointerA)
	*pointerA = 20
	println(a)
	address := &a
	println(address)
	println(*address)
}
