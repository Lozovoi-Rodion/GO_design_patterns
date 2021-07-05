package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/builder/examples"
	"strings"
)

func main() {
	hello := "hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Println(sb.String())

	b := builder.NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")
	b.AddChildFluent("li", ", fluently").AddChildFluent("li", "and elegant")
	fmt.Println(b.String())

	pb := builder.NewPersonBuilder{}
}
