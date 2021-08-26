package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/memento/examples"
)

// Memento - is a behavioral dp that lets you save and restore the previous state of an object
// without revealing the details of its implementation

func main() {
	ba, m := memento.NewBankAccount(100)

	m1 := ba.Deposit(50)
	fmt.Println(m1)
	fmt.Println(ba)
	ba.Restore(m)
	fmt.Println(ba)
}
