package main

import (
	"fmt"
)


// Composite:
// Composition lets us make compound objects (through embedding) if object A need to use object's B fields/methods
// Composite - is a DP that allows treating singular(scalar) objects and compositions of objects in a uniform manner
// Summary:
// -
// -
// -

func main()  {
	drawing := composite.GraphicObject{"My drawing", "", nil}
	drawing.Children = append(drawing.Children, *composite.NewCircle("Red"))
	drawing.Children = append(drawing.Children, *composite.NewSquare("Yellow"))

	group := composite.GraphicObject{"Group1", "", nil}
	group.Children = append(group.Children, *composite.NewCircle("Blue"))
	group.Children = append(group.Children, *composite.NewSquare("Blue"))

	drawing.Children = append(drawing.Children, group)

	fmt.Println(drawing.String())
}