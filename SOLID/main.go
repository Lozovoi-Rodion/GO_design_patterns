package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/SOLID/LSP"
	"github.com/Lozovoi-Rodion/GO_design_patterns/SOLID/OCP"
	"github.com/Lozovoi-Rodion/GO_design_patterns/SOLID/SRP"
)

func main() {
	/*SRP:*/
	j := SRP.Journal{}
	j.AddEntry("This is my first entry")
	j.AddEntry("This is my last entry")
	fmt.Println(j.String())

	/*OCP*/
	apple := OCP.Product{"Apple", OCP.Green, OCP.Small}
	tree := OCP.Product{"Tree", OCP.Green, OCP.Large}
	car := OCP.Product{"Car", OCP.Red, OCP.Medium}
	products := []OCP.Product{apple, tree, car}
	fmt.Printf("Green products (old):\n")
	f := OCP.Filter{}
	for _, v := range f.FilterByColor(products, OCP.Green) {
		fmt.Printf(" - %s, is green \n", v.Name)
	}

	fmt.Printf("Green products (new):\n")
	bf := OCP.BetterFilter{}
	greenSpec := OCP.ColorSpecification{OCP.Green}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s, is green \n", v.Name)
	}

	fmt.Printf("Green and large products:\n")
	largeSpec := OCP.SizeSpecification{OCP.Large}
	lgSpec := OCP.AndSpecification{greenSpec, largeSpec}
	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf(" - %s is green and large \n", v.Name)
	}

	/*LSP*/
	rc := &LSP.Rectangle{2, 3}
	LSP.UseIt(rc)

	sq := LSP.NewSquare(5)
	LSP.UseIt(sq)
}
