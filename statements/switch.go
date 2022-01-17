package main

import (
	"fmt"
	"time"
)

func getMacVersion() int {
	return 10
}

func main() {
	switch ver := getMacVersion(); ver {
	case 11:
		fmt.Println("Big sur")
	case 12:
		fmt.Println("Monterey")
	default:
		fmt.Println("Old version", ver)
	}

	t := time.Now()
	fmt.Println(t.Hour())

	switch {
	case t.Hour() < 12:
		fmt.Println("Beforenoon")
	case t.Hour() > 12:
		fmt.Println("Afternoon")
	}
}
