package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/observer/examples"
)

// Motivation:
// - be informed when certain things happen (ob's field changes, ob does smth, external event occurs)
// - we want to listen to events and be notified when they occur
// - 2 participants: observable, observer
// An observer is an object that wishes to be informed about events happening in the system.
// The entity generating the events is an observable
// Summary:
// - observer is an intrusive approach
// - must provide a way of clients to subscribe
// -event data sent from observable to all observers
// data represented as interface{}
// unsubscription is possible

func main() {
	p := observer.NewPerson("Sanya")
	ds := observer.NewDoctorService()
	p.Subscribe(ds)

	p.FireSubs()

	p2 := observer.NewPerson("Dimon")
	p2.SetAge(13)
	tm := &observer.TrafficManagement{p2.Observable}
	er := &observer.ElectoralRoll{}
	p2.Subscribe(tm)
	p2.Subscribe(er)

	for i := 14; i < 20; i++ {
		fmt.Println("Setting the age to", i)
		p2.SetAge(i)
	}
}
