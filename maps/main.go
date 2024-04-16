package main

import "fmt"

func main() {
	salaries := map[string]int{
		"jose":   2300,
		"duarte": 200,
	}

	fmt.Println(salaries["jose"])

	delete(salaries, "jose")

	fmt.Println(salaries["jose"])

	salaries["neto"] = 5000

	fmt.Println(salaries["neto"])

	days := make(map[string]int)
	week := map[string]int{}
	days["Monday"] = 1
	days["Tuesday"] = 2
	days["wednesday"] = 3
	days["Thursday"] = 4
	days["Friday"] = 5
	days["Saturday"] = 6
	days["Sunday"] = 7

	fmt.Println(days)
	fmt.Println(week)

	for name, salary := range salaries {
		fmt.Printf("%s's salary is %d\n", name, salary)
	}

	for _, salary := range salaries {
		fmt.Printf("salary is %d\n", salary)
	}

}
