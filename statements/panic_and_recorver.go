package main

import "fmt"

func thirdPartyConnectDatabase() {
	panic("can't connect to the database")

}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDatabase()
}

func main() {
	save()
	fmt.Println("OK?")
}
