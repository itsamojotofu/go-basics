package main

import "fmt"

func main() {
	b := []byte{72, 73}
	fmt.Println(b)
	fmt.Println(string(b))

	y := []byte("HI")
	fmt.Println(y)
	fmt.Println(string(y))
}
