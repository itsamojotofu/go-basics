package main

import "fmt"

func main() {
	l := []string{"go", "c", "typescript"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	for i, v := range l {
		fmt.Println(i, v)
	}

	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 80}

	for item, price := range m {
		fmt.Println(item, price)
	}
	for item := range m {
		fmt.Println(item)
	}
	for _, price := range m {
		fmt.Println(price)
	}
}
