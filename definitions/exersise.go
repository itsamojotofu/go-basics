package main

import "fmt"

func main() {
	pi := 3.14
	fmt.Printf("%T %v\n", pi, pi)

	s := []int{1, 2, 5, 6, 2, 3, 1}
	// output is [5 6]
	fmt.Println(s[2:4])

	m := map[string]int{
		"Apple":  20,
		"Banila": 40,
		"Cherry": 30,
	}

	fmt.Printf("%T %v\n", m, m)
}
