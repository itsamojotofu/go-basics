package main

import "fmt"

func cal(price, item int) (result int) {
	result = price * item
	return result
}

func convert(price int) float64 {
	return float64(price)
}

func main() {
	total := cal(100, 4)
	fmt.Println(total)

	f := func(x float64) {
		fmt.Println("total price", x)
	}
	f(convert(total))
}
