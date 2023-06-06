package main

import "fmt"

func main() {
	//var slice []int
	slice := make([]int, 0, 5)

	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println(slice[0])
}
