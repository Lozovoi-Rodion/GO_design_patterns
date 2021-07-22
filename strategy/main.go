package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/strategy/examples"
)

// Strategy - separates an algorithm into its `skeleton` and concrete implementation
// steps, which can be varied at run-time
// Summary:
// - define an algorithm at a high level
// - define the interface you expect strategy to follow
// - support the injection of the strategy into the high-level algorithm

func main() {
	tp := strategy.NewTextProcessor(&strategy.MarkdownListStrategy{})
	tp.AppendList([]string{"This", "is", "markdown", "list"})
	fmt.Println(tp.String())

	tp.Reset()
	tp.SetOutputFormat(strategy.Html)
	tp.AppendList([]string{"This", "is", "html", "list"})
	fmt.Println(tp.String())
}