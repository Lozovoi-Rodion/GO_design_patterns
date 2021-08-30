package main

// state - is a dp in which the object's behavior determined by its state.
// An object transitions from one state to another ( something need to trigger
// a transition ).
// A formalized construct which manages state and transitions is called
// a state machine

func main()  {
	sw := state.NewSwitch()
	sw.On()
}
