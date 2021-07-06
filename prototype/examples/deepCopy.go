package prototype

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name string
	Address *Address
}

// DeepCopyExample can not be scaled
func DeepCopyExample()  {
	john := Person{"John", &Address{"Johnny Avenue 13", "London", "UK"}}
	jane := john
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"

	fmt.Println(john.Name, john.Address)
	fmt.Println(jane.Name, jane. Address)
}