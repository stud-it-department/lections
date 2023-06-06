package main

import (
	"fmt"
	"main/internal/sort"
)

func main() {
	data := []int{5, 6, 3, 1, 5, 7, 8, 3, 5, 2, 9, 0}

	fmt.Println(data)
	
	err := sort.Sort(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)
}
