package main

import "fmt"

func main() {
	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	n = append(n, 1, 2, 3, 4, 5, 6, 7)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	x := make([]int, 0)
	var y []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(x), cap(x), x)
	fmt.Printf("len=%d cap=%d value=%v\n", len(y), cap(y), y)

	// examine the differencs between len and cap
	exam_1 := make([]int, 5)
	for i := 0; i < 5; i++ {
		exam_1 = append(exam_1, i)
		fmt.Println(exam_1)
	}
	fmt.Println(exam_1)

	exam_2 := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		exam_2 = append(exam_2, i)
		fmt.Println(exam_2)
	}
	fmt.Println(exam_2)
}
