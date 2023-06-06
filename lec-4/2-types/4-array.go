package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr := [10]int{5, 4, 3, 2, 1}

	fmt.Println(unsafe.Sizeof(arr), arr)
	fmt.Printf("%p [%p, %p]\n", &arr, &arr[0], &arr[9])
}
