package main

import "fmt"

func main() {
	// var t,f bool = true, false
	t, f := true, false
	fmt.Printf("%T %v\n", t, t)
	fmt.Printf("%T %v\n", f, f)

	fmt.Println(true && true)
	fmt.Println(true && false)

	fmt.Println(!true)
	fmt.Println(!false)
}
