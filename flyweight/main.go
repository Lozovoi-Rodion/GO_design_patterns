package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/flyweight/examples"
)

// Motivation : avoid redundancy when storing data
// Flyweight - is a space optimization technique that lets us use less memory by storing externally
// the data associated with similar object

func main()  {
	text := "This is a brave new world"
	ft := flyweight.NewFormattedText(text)
	ft.Capitalize(10,15)
	fmt.Println(ft.String())
}
