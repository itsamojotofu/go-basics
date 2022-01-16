package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		} else if i == 8 {
			break
		}
		fmt.Println(i)
	}

	db := 1
	for db < 30 {
		db += db
		fmt.Println(db)
	}

	// if you apply no conditions under "for" statement, it will cause an infinite loop
	/*
		for {
			fmt.Println("HIIT!")
		}
	*/
}
