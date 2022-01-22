package main

import (
	"fmt"

	".packages/mylib"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))
}
