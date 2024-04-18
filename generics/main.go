package main

type N2 int

type N interface {
	~int | ~float64
}

func main() {
	i := map[string]int{
		"Jose":   123,
		"Duarte": 234,
		"M":      346,
		"N":      456,
	}
	f := map[string]float64{
		"Jose":   123.1,
		"Duarte": 234.2,
		"M":      346.3,
		"N":      456.4,
	}
	z := map[string]N2{
		"Jose":   1,
		"Duarte": 2,
		"M":      3,
		"N":      4,
	}

	s1 := Sum(i)
	s2 := Sum(f)
	s3 := Sum(z)
	println(s1, s2, s3)
	println(Compare(10, 10))
	println(Compare(10, 11.0))
}

// func Sum[T int | float64](m map[string]T) T {
// 	var sum T
// 	for _, v := range m {
// 		sum += v
// 	}
// 	return sum
// }

// constraint
func Sum[T N](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func Compare[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}
