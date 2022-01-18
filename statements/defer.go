package main

import (
	"fmt"
	"os"
)

func main() {
	/* 	defer fmt.Println("Noodles")
	fmt.Println("Ramen") */

	/* 	fmt.Println("Start")
	   	defer fmt.Println(1)
	   	defer fmt.Println(2)
	   	defer fmt.Println(3)
	   	fmt.Println("End") */

	// defer also used when we open a file and have no error
	file, _ := os.Open("../hello_world.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}
