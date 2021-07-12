package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/flyweight/examples"
)

// Motivation : avoid redundancy when storing data
// Flyweight - is a space optimization technique that lets us use less memory by storing externally
// the data associated with similar object
// Summary:
// - store common data externally
// - specify an index or a pointer into the external data store
// - define the idea of 'ranges' on homogeneous collections and store data related to those ranges
func main()  {
	// inefficient approach
	text := "This is a brave new world"
	ft := flyweight.NewFormattedText(text)
	ft.Capitalize(10,15)
	fmt.Println(ft.String())

	// correct approach
	bft := flyweight.NewBetterFormattedText(text)
	bft.Range(16,19).Capitalize = true
	fmt.Println(bft.String())

	// second example

	// not optimized
	john := flyweight.NewUser("John Doe")
	jane := flyweight.NewUser("Jane Doe")
	alsoJane := flyweight.NewUser("Jane Smith")
	fmt.Println(john.FullName)
	fmt.Println(jane.FullName)
	fmt.Println("Memory taken by users:",
		len([]byte(john.FullName)) +
			len([]byte(alsoJane.FullName)) +
			len([]byte(jane.FullName)))

	// optimized with flyweight usage
	john2 := flyweight.NewUser2("John Doe")
	jane2 := flyweight.NewUser2("Jane Doe")
	alsoJane2 := flyweight.NewUser2("Jane Smith")
	fmt.Println(john2.FullName())
	fmt.Println(jane2.FullName())
	totalMem := 0
	for _, a := range flyweight.AllNames {
		totalMem += len([]byte(a))
	}
	totalMem += len(john2.Names)
	totalMem += len(jane2.Names)
	totalMem += len(alsoJane2.Names)
	fmt.Println("Memory taken by users2:", totalMem)
}
