package factory

import "fmt"

type iPerson interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s. I am %d years old \n", p.name, p.age)
}

func (tp *tiredPerson) SayHello() {
	fmt.Printf("Hi, I am tired %s. I am %d years old \n", tp.name, tp.age)
}

func NewPerson2(name string, age int) iPerson {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}
