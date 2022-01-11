package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n[3:5])
	fmt.Println(n[:2])
	fmt.Println(n[:])

	// create two layers in array
	double := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(double)

	double = append(double, []int{7, 8, 9})

	fmt.Println(double)

}
