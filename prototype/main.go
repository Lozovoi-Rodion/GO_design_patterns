package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/prototype/examples"
)

//Summary:
// - to implement a prototype, partially construct an object and store it somewhere
// - deep copy the prototype
// - customize  the resulting instance
// - A prototype factory provides a convenient API for using prototypes

func main() {
	prototype.DeepCopyExample()

	john := prototype.Person{"John", &prototype.Address{"Starway Road 32", "Bejing", "China"}, []string{"Li", "Lu"}}
	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Kuan")

	fmt.Println("---!!!!---")
	fmt.Println(john.Name, john.Address, john.Friends)
	fmt.Println(jane.Name, jane.Address, jane.Friends)

	jane2 := john.DeepSerializedCopy()
	jane2.Name = "Jane"
	jane2.Address.StreetAddress = "321 Baker St"
	jane2.Friends = append(jane2.Friends, "Kuan")

	fmt.Println("---!!!!---")
	fmt.Println(john.Name, john.Address, john.Friends)
	fmt.Println(jane.Name, jane.Address, jane.Friends)
	fmt.Println(jane2.Name, jane2.Address, jane2.Friends)

	daniel := prototype.NewMainOfficeEmployee("Daniel", 1)
	garry := prototype.NewAuxOfficeEmployee("Garry", 3)

	fmt.Println(daniel)
	fmt.Println(garry)
}
