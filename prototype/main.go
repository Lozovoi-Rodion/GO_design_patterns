package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/prototype/examples"
)

func main() {
	prototype.DeepCopyExample()

	john := prototype.Person{"John", &prototype.Address{"Starway Road 32", "Bejing", "China"}, []string{"Li", "Lu"}}
	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Kuan")

	fmt.Println("---!!!!---")
	fmt.Println(john.Name, john.Address, john.Friends)
	fmt.Println(jane.Name, john.Address, jane.Friends)
}
