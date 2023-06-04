package main

import "fmt"

type IAnimal interface {
	MakeSound() string
}

type Dog struct {
	Age      int
	GavCount int
}

func (d *Dog) Gav() string {
	return "gav gav"
}

func (d *Dog) MakeSound() string {
	return d.Gav()
}

// --------

type Cat struct {
	Age       int
	MeowCount int
}

func (c *Cat) Meow() string {
	return "meow"
}

func (c *Cat) MakeSound() string {
	return c.Meow()
}

func main() {
	d := &Dog{}
	c := &Cat{}

	// 	c := Cat{} - неправильно! потому что интерфейс - указатель на что-то!

	GetSound(d)
	GetSound(c)
}

func GetSound(a IAnimal) {
	fmt.Println(a.MakeSound())
}
