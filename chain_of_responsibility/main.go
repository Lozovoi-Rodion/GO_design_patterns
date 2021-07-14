package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/chain_of_responsibility/examples"
	"sync"
)

// command query separation (CQS)
// Command = asking for an action or change (e.g. please set your attack value to 2
// Query = asking for information (e.g. please give me attack value
// CQS - having separate means of sending commands and queries to (e.g., direct field values)

// Summary:
// - CoR can be implemented as a linked list of pointers or a centralized construct
// - Enlist objects in the chain controlling their order
// - Control object removed from chain

func main() {
	murlock := chain_of_responsibility.NewCreature("Murlock", 2, 1)
	fmt.Println(murlock.String())

	root := chain_of_responsibility.NewCreatureModifier(murlock)

	//root.Add(chain_of_responsibility.NewSilenceModifier(murlock))

	root.Add(chain_of_responsibility.NewDoubleAttackModifier(murlock))
	root.Add(chain_of_responsibility.NewIncreasedDefenseModifier(murlock))
	root.Add(chain_of_responsibility.NewDoubleAttackModifier(murlock))
	root.Handle()
	fmt.Println(murlock.String())

	//broker chain:
	game := &chain_of_responsibility.Game{sync.Map{}}
	goblin := chain_of_responsibility.NewCreature2(game, "Superior Goblin", 2, 3)
	fmt.Println(goblin.String())

	m := chain_of_responsibility.NewDoubleAttackModifier2(game, goblin)
	fmt.Println(goblin.String())
	m.Close()
	fmt.Println(goblin.String())
}
