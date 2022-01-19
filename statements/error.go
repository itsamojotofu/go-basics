package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./range.go")
	if err != nil {
		log.Fatalln("Error!")
	}
	defer file.Close()
	data := make([]byte, 100)
	// initializer ":=" can be reused if at least one variable is new.
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))
}
