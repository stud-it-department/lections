package main

import (
	"fmt"
	"os"
)

func main() {
	fd, err := os.OpenFile("lec-4/4-linux/file", os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	// RAII
	defer func() {
		_ = fd.Close()
	}()

	if _, err := fd.WriteString("hello, world!"); err != nil {
		fmt.Println(err)
		return
	}
}
