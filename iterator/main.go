package main

import (
	"fmt"
	iterator "github.com/Lozovoi-Rodion/GO_design_patterns/iterator/examples"
)

// Summary:
// - an iterator specifies how can you traverse an object
// - moves along the iterated collection, indicating when last element has been reached
// - not idiomatic in go

func main() {
	p := iterator.Person{"Alexander", "Graham", "Bell"}
	for _, name := range p.Names() {
		fmt.Println(name)
	}

	p2 := iterator.Person{"Nikola", "", "Tesla"}
	for name := range p2.NamesGenerator() {
		fmt.Println(name)
	}

	for it := iterator.NewPersonNameIterator(&p); it.MoveNext(); {
		fmt.Println(it.Value())
	}

	root := iterator.NewNode(1, iterator.NewNode(2, iterator.NewTerminalNode(21), iterator.NewTerminalNode(22)), iterator.NewTerminalNode(3))
	it := iterator.NewInOrderIterator(root)

	for it.MoveNext() {
		fmt.Printf("%d", it.Current.Value)
	}
	fmt.Printf("\n")

	t := iterator.NewBinaryTree(root)
	for i := t.InOrder(); i.MoveNext(); {
		fmt.Printf("%d,", i.Current.Value)
	}
	fmt.Printf("\n")
}
