package main

// Mediator - a component that facilitates communication between other components without
// them necessarily being aware of each other or having direct (reference) access to each other


func main()  {
	room := mediator.NewChatRoom()

	john := mediator.NewPerson("John")
	jane := mediator.NewPerson("Jane")

	room.Join(john)
	room.Join(jane)

	john.Say("hi everyone")
	jane.Say("hi John")
}
