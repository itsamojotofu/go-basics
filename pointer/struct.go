package main

import "fmt"

type Vertex struct {
	X, Y int
	S    string
}

func UpdateX(v Vertex) {
	v.X = 100
}

func UpdateX2(v *Vertex) {
	v.X = 200

}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)

	v1 := Vertex{2, 4, "6"}
	fmt.Println(v1)

	vn := Vertex{}
	fmt.Println(vn)

	fmt.Println("-----")

	// looking into the innerworkings of pointer
	p1 := Vertex{1, 2, "pointer doesn't work..."}
	UpdateX(p1)
	fmt.Println(p1)

	p2 := &Vertex{1, 2, "pointer works!!"}
	UpdateX2(p2)
	fmt.Println(p2)
}
