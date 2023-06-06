package main

import "fmt"

func main() {
	slice1 := []int{1, 1, 2, 3}

	slice1 = append(slice1, 5, 7)

	slice2 := []int{9, 11}
	slice2 = append(slice1, slice2...)

	fmt.Println(slice1, slice2)
}
