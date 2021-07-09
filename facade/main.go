package main

import "github.com/Lozovoi-Rodion/GO_design_patterns/facade/examples"
// Facade provides a simple, easy to understand/user interface over
// a large and sophisticated body of code


func main()  {
	c := facade.NewConsole()
	u := c.GetCharacterAt(1)
}