package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 2
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 2.4
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var s string = "12"
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v %d\n", i, i, i)

	h := "hello"
	fmt.Println(string(h[4]))
}
