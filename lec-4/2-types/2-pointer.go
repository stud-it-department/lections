package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 10
	b := 20

	ptr := &a

	fmt.Printf("Address of a=%d: %p\n", a, &a)
	fmt.Printf("Address of ptr=&i=%d: %p\n", *ptr, ptr)
	fmt.Printf("Address of pointer ptr=&i=%d: %p\n", *ptr, &ptr)

	fmt.Println(*ptr)

	*ptr = 17
	fmt.Println(a)

	ptr = &b
	fmt.Println(*ptr)

	fmt.Println(unsafe.Sizeof(ptr))
}
