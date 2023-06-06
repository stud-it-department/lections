package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["apple"] = 3

	ok := m["banana"]
	fmt.Println(ok)

	if _, ok := m["banana"]; !ok {
		fmt.Println("banana does not exists")
	}

	if appleCount, ok := m["apple"]; !ok {
		fmt.Println("apple does not exists")
	} else {
		fmt.Println("apple exists")
		fmt.Println(appleCount)
	}
}
