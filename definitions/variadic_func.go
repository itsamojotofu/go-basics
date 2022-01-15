package main

import "fmt"

func steady(param1, param2 int) {
	fmt.Println(param1, "and", param2)
}

func variadic(params ...int) {
	fmt.Println(len(params), params)
}

func main() {
	steady(10, 20)
	variadic(10)
	variadic(10, 20)
	variadic(10, 20, 30)
	variadic()
}
