package main

import (
	"fmt"

	"packages/mylib"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	// this relative path must be replaced with any absolute path
	fmt.Println(mylib.Average(s))
}
