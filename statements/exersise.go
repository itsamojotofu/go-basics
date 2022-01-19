package main

import "fmt"

func main() {
	min := 0
	l := []int{100, 300, 50, 20, 50, 3, 5, 10, 5}
	for i, v := range l {
		if i == 0 {
			min = v
			continue
		}

		if v < min {
			min = v
		}
	}

	fmt.Println(min)

	m := map[string]int{
		"apple":   300,
		"banana":  200,
		"cabbage": 180,
		"dragon":  800,
		"egg":     200,
		"fanta":   120,
	}

	sum := 0

	for _, price := range m {
		sum += price
	}

	fmt.Println(sum)

}
