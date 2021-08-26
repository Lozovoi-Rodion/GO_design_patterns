package main

import "github.com/Lozovoi-Rodion/GO_design_patterns/mediator/examples"

// Mediator - a component that facilitates communication between other components without
// them necessarily being aware of each other or having direct (reference) access to each other


func main()  {
	room := mediator.NewChatRoom()

	john := mediator.NewPerson("John")
	jane := mediator.NewPerson("Jane")

	room.Join(jane)
	room.Join(john)

	john.Say("hi everyone")
	jane.Say("hi John")

	simon := mediator.NewPerson("Simon")
	room.Join(simon)

	jane.PrivateMessage("Simon", "hello, Semen")
}
