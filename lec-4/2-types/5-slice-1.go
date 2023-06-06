package main

import (
	"fmt"
	"unsafe"
)

func main() {
	slice := []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	fmt.Println(unsafe.Sizeof(slice), slice)
	fmt.Printf("%p [%p - %p]\n", &slice, &slice[0], &slice[9])
}
