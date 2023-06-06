package main

import (
	"fmt"
)

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Printf("Address of arr: %p\n", &arr)

	slice := arr[:]
	fmt.Printf("Address of slice: %p\n", &slice)

	fmt.Printf("Address of the slice underlying array: %p\n", slice)

	fmt.Printf("Address of 0th element of the slice: %p\n", &slice[0])
}
