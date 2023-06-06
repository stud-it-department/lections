package main

import "time"

func main() {
	x := getValue()
	time.Sleep(time.Second * 5)
	println(*x / 2)
}

func getValue() *int {
	res := 10
	return &res
}
