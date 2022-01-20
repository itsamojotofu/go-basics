package main

import "fmt"

func main() {
	var n int = 1
	fmt.Println(n)
	fmt.Println(&n)

	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)
}
