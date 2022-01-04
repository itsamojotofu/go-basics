package main

import "fmt"

// var_dec can be made outside of main func like the following

var (
	i    int     = 1
	pai  float64 = 3.14
	s    string  = "test"
	t, f bool    = true, false
)

func main() {
	// no need to declare type when using ':='
	i2 := 1
	pai2 := 3.14
	s2 := "test"
	t2, f2 := true, false
	fmt.Println(i, pai, s, t, f)
	fmt.Println(i2, pai2, s2, t2, f2)
	fmt.Printf("%T\n", pai2)
	fmt.Printf("%T\n", s2)
}
