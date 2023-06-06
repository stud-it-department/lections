package main

import (
	"fmt"
	"strings"
)

func main() {
	example := "hello world!"

	for _, runa := range example {
		fmt.Printf("%t\n", runa)
		// rune() - int32 потому что там так же могут быть символы дичь - например китайские
	}

	rplc := strings.ReplaceAll(example, "!", "?")

	fmt.Println(rplc)
}
