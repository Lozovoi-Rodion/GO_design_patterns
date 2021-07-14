package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/command/examples"
)
// Command - dp that represents an instruction to perform a particular action.
// Contains all the information necessary for the action to be taken
// Summary:
// - encapsulate all the details of an operation in a separate object
// - define f-ns for applying the command itself or elsewhere
// - optionally define instructions for undying the command
// - can create composite commands (a.k.a. macros) + CoR + Composite


func main()  {
	ba := command.NewBankAccount(0)
	cmd := command.NewBankAccountCommand(ba, command.Deposit, 300)
	cmd.Call()
	fmt.Println(ba)
	cmd2 := command.NewBankAccountCommand(ba, command.Withdraw, 500)
	cmd2.Call()
	fmt.Println(ba)
	cmd2.Undo()
	fmt.Println(ba)

	from := command.NewBankAccount(100)
	to := command.NewBankAccount(0)
	mtc := command.NewMoneyTransferCommand(from, to, 80) // try 1000
	mtc.Call()

	fmt.Println("from=", from, "to=", to)

	fmt.Println("Undoing...")
	mtc.Undo()
	fmt.Println("from=", from, "to=", to)
}
