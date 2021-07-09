package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/composite/examples"
)


// Composite:
// Composition lets us make compound objects (through embedding) if object A need to use object's B fields/methods
// Composite - is a DP that allows treating singular(scalar) objects and compositions of objects in a uniform manner
// Summary:
// - objects can use other objects via composition
// - some composed and singular objects need similar/identical behaviour
// - composite dp lets us treat both types of objects uniformly
// - iteration supported via iterator dp

func main()  {
	drawing := composite.GraphicObject{"My drawing", "", nil}
	drawing.Children = append(drawing.Children, *composite.NewCircle("Red"))
	drawing.Children = append(drawing.Children, *composite.NewSquare("Yellow"))

	group := composite.GraphicObject{"Group1", "", nil}
	group.Children = append(group.Children, *composite.NewCircle("Blue"))
	group.Children = append(group.Children, *composite.NewSquare("Blue"))

	drawing.Children = append(drawing.Children, group)

	fmt.Println(drawing.String())

	// Neuron example
	neuron1, neuron2 := &composite.Neuron{}, &composite.Neuron{}
	layer1, layer2 := &composite.NeuronLayer{}, &composite.NeuronLayer{}

	composite.Connect(neuron1, neuron2)
	composite.Connect(neuron1, layer1)
	composite.Connect(neuron1, layer1)
	composite.Connect(layer2, neuron1)
}