package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "grape": 300}
	fmt.Println(m)
	m["peach"] = 250
	fmt.Println(m)
	fmt.Println(m["peach"])
	fmt.Println(m["nothing"])
	v, ok := m["apple"]
	fmt.Println(v, ok)
	v, ok = m["nothing"]
	fmt.Println(v, ok)

	m2 := make(map[int]string)
	fmt.Println(m2)
	m2[1] = "January"
	fmt.Println(m2)

	// if we use var to declare map or slice, it will be nil

	// var m3 map[int]string
	// m3[1] = "January"
	// fmt.Println(m3)
}
