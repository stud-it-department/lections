package main

import (
	"fmt"
	"unsafe"
)

type Vertex struct {
	X, Y int
}

func main() {
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)

	v1.Y = 5e5
	p.X = 1e9

	fmt.Println(unsafe.Sizeof(v1), unsafe.Sizeof(p), v1, v2, v3, p)
}
