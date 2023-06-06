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
	c := Cat{}

	// interface = any

	GetSound(d)
	GetSound(&c) // получаем указатель этого типа

	Ex(d)
	Ex(12)
}

func GetSound(a IAnimal) {
	fmt.Println(a.MakeSound())
}

func Ex(a interface{}) {
	if _, ok := a.(IAnimal); !ok {
		fmt.Println("not ok!")
		return
	}

	fmt.Println("ok!")
}
