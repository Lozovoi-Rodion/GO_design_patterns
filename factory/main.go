package main

import (
	"fmt"
)


func main() {
	p := personFactory.NewPerson("John", 22)
	p.EyeCount = 4

	fmt.Println(p)
}
