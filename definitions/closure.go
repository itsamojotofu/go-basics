package main

import "fmt"

// by taking out this function outside the main func, x would be overwritten by other declarations
func increment() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func main() {
	counting := increment()
	fmt.Println(counting())
	fmt.Println(counting())
	fmt.Println(counting())

	circle := circleArea(3.14)
	fmt.Println(circle(5))
}
