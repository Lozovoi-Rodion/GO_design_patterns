package observer

import (
	"container/list"
	"fmt"
)

type Observable struct {
	subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
	for s := o.subs.Front(); s != nil; s = s.Next() {
		if s.Value.(Observer) == x {
			o.subs.Remove(s)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for s := o.subs.Front(); s != nil; s = s.Next() {
		s.Value.(Observer).Notify(data)
	}
}

func (o *Observable) FireSub(x Observer, data interface{}) {
	for s := o.subs.Front(); s != nil; s = s.Next() {
		if s.Value.(Observer) == x {
			s.Value.(Observer).Notify(data)
		}
	}
}

type Observer interface {
	Notify(data interface{})
}

type PropertyChange struct {
	Name  string // name of changed property
	Value interface{}
}

type Person struct {
	Observable
	Name string
	age  int
}

func NewPerson(name string) *Person {
	return &Person{Observable: Observable{new(list.List)}, Name: name}
}

func (p *Person) Age() int {
	return p.age
}

// property observer
func (p *Person) SetAge(age int) {
	if p.age == age {
		return
	}

	oldCanVote := p.CanVote()
	p.age = age
	p.Fire(PropertyChange{"Age", p.age})

	if oldCanVote != p.CanVote() {
		p.Fire(PropertyChange{"CanVote", p.CanVote()})
	}
}

func (p *Person) CatchACold(x Observer) {
	p.FireSub(x, p.Name)
}

func (p *Person) CleanTheHouse(x Observer) {
	p.FireSub(x, p.Name)
}

func (p *Person) CanVote() bool {
	return p.age >= 18
}

type ElectoralRoll struct {
}

func (e ElectoralRoll) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Congratulations, you can perform your civil duty!")
		}
	}
}

func (p *Person) FireSubs() {
	p.Fire(p.Name)
}

type DoctorService struct{}

func NewDoctorService() *DoctorService {
	return &DoctorService{}
}

func (d *DoctorService) Notify(data interface{}) {
	fmt.Printf("A doctor has been called for %s \n", data.(string))
}

type TrafficManagement struct {
	O Observable
}

func (t *TrafficManagement) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Value.(int) >= 16 {
			fmt.Printf("You are allowed to drive. Your %s is %d \n", pc.Name, pc.Value)
			t.O.Unsubscribe(t)
		}
	}
}
