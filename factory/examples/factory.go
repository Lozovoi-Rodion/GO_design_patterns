// Package factory: simple factory example
package factory

type Person1 struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson1(name string, age int) *Person1 {
	return &Person1{name, age, 2}
}