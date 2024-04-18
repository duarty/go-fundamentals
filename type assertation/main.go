package main

func main() {

	var v interface{} = "hello world"
	println(v.(string))

	res, ok := v.(int)
	println(res, ok)

	// res2 := v.(int)
	// println(res2)
}
