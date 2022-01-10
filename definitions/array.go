package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b [2]int = [2]int{1, 2}
	// below causes error because array cannot be resized
	// b = append(b, 3)
	fmt.Println(b)

	var c []int = []int{10, 20}
	c = append(c, 30)
	fmt.Println(c)

}
