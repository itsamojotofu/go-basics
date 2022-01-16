package main

import "fmt"

func even_check(n int) string {
	if n%2 == 0 {
		return "yes"
	} else {
		return "no"
	}
}

func main() {
	num := 12
	if num%2 == 0 {
		fmt.Println("even")
	} else if num%3 == 0 {
		fmt.Println("by three")
	} else {
		fmt.Println("others")
	}

	if result := even_check(18); result == "yes" {
		fmt.Println("even!")
	}
}
