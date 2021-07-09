package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/facade/examples"
)
// Facade provides a simple, easy to understand/user interface over
// a large and sophisticated body of code
// Summary:
// - facade provides simplified API over a set of components
// - may wish to expose (optionally) internals through the facade
// - may allow users to use more complex API if they need to

func main()  {
	c := facade.NewConsole()
	u := c.GetCharacterAt(1)
	fmt.Println(u)
}