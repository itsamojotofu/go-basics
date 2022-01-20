package main

import "fmt"

func main() {
	var p *int = new(int)
	var p2 *int

	fmt.Println(p, p2)
	fmt.Println(*p)
	*p++
	fmt.Println(*p)
	// runtime error: invalid memory address or nil pointer dereference
	// fmt.Println(*p2)

	s := make([]int, 0)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", p)
}
