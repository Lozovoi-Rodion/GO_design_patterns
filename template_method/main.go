package main

import 	"github.com/Lozovoi-Rodion/GO_design_patterns/template_method/examples"

// Overview:
// algorithms can be decomposed into common parts + specifics
// Strategy pattern does this through composition:
// - high-level algo uses an interface
// - concrete implementation implements the interface
// - we keep a pointer to the interface; provide concrete implementation
// Template Method performs a similar operation, but:
// it's typically a function, not a struct with a reference to the implementation
// can still use interfaces (just like Strategy): or can be functional (takes several fns as params)
// Template Method - is a dp where a skeleton defined in a function. Function can either use an interface (like Strategy)
// or can take several fns as args.
// Summary:
// - define interface with common operations
// - make use of those operations inside a function

func main() {
	chess := template_method.NewGameOfChess()
	template_method.PlayGame(chess)


}
