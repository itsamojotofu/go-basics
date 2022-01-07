package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("Hello" + "World")
	fmt.Println("Hello world"[1])

	fmt.Println(strings.Replace("Hello World", "l", "k", 2))
	fmt.Println(strings.Contains("Hello World", "World"))
	fmt.Println(`One
	Two
		Three`)
	// How to show double quatations as a string
	fmt.Println(`"Fantastic"`)
	fmt.Println("\"Fantastic\"")
}
