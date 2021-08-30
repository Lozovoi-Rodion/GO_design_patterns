package main

import (
	"fmt"
)

// Memento - is a behavioral dp that lets you save and restore the previous state of an object
// without revealing the details of its implementation
// Difference between Flyweight:
// 1. Memento is used only to be fed back into the system:
// - no public/mutable state
// - no methods
// 2. A flyweight is similar to ordinary reference to an object
// - can mutate state
// - can provide its own functionality(fields,methods)
// Summary:
// - used to roll back state arbitrarily
// - can be used to implement undo/redo
// - doesn't have to expose its own state
// - that's simple token/handle with no methods of its own

func main() {
	ba := memento.NewBankAccount(100)
	ba.Deposit(50)
	ba.Deposit(25)
	m := ba.Deposit(25)
	fmt.Println(ba)

	ba.Undo()
	fmt.Println("Undo 1: ", ba)
	ba.Undo()
	fmt.Println("Undo 2: ", ba)
	ba.Redo()
	fmt.Println("Redo 1: ", ba)
	ba.Undo()
	fmt.Println("Undo 3: ", ba)
	ba.Undo()
	fmt.Println("Undo 4: ", ba)
	fmt.Println("m1: ", m)
	ba.Restore(m)
	fmt.Println("Restore 1: ", ba)
	ba.Undo()
	fmt.Println("Undo 5: ", ba)
}
