package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/builder/examples"
)

//Summary:
// - a builder is a separate component used for building ab object
// - it provides API for succinct construction if object is complicated
// - to make builder fluent - return the receiver. It allows chaining
// - different facets of an object can be built with different builder types working
// in tandem via a common struct

func main() {
	b := builder.NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")
	b.AddChildFluent("li", ", fluently").AddChildFluent("li", "and elegant")
	fmt.Println(b.String())

	pb := builder.NewPersonBuilder()
	pb.
		Lives().
		At("Golandish Road").
		In("Kyiv").
		WithPostcode("KGL777").
		Works().
		At("ITishka").
		AsA("engineer").
		Earning(99999)
	person := pb.Build()
	fmt.Println(*person)

	builder.SendEmail(func(b *builder.EmailBuilder) {
		b.From("foo@bar.com").
			To("bar@baz.com").
			Subject("Meeting").
			Body("Hello, do you want to meet?")
	})

	pfb := builder.PersonFpBuilder{}
	pf := pfb.Called("Rodya").WorksAsA("dev").Build()
	fmt.Println(*pf)
}
