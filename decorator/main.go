package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/decorator/examples"
)

// - used to augment an object with additional functionality
// - do not want to rewrite or alter existing code (OCP)
// - want to keep new functionality separate (SRP)
// - need to be able to interact with existing structs
// Decorator - is a dp that facilitates the addition of behaviors to individual objects through embedding

// Summary:
// - decorator embeds decorated objects
// - adds utility fields and methods to augment the object's features
// - often used to emulate inheritance (may require extra work)

func main() {
	d := decorator.NewDragon()
	d.SetAge(10)
	d.Fly()
	d.Crawl()

	circle := decorator.Circle{2}
	fmt.Println(circle.Render())

	redCircle := decorator.ColoredShape{&circle, "red"}
	fmt.Println(redCircle.Render())

	rhsCircle := decorator.TransparentShape{&redCircle, 0.5}
	fmt.Println(rhsCircle.Render())
}
