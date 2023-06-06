package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a bool
	var str string = "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
	var i int
	var ui uint
	var uip uintptr
	var b byte
	var r rune

	fmt.Println(unsafe.Sizeof(a),
		unsafe.Sizeof(str),
		unsafe.Sizeof(i),
		unsafe.Sizeof(ui),
		unsafe.Sizeof(uip),
		unsafe.Sizeof(b),
		unsafe.Sizeof(r))
}
