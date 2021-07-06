package prototype

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{a.StreetAddress, a.City, a.Country}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

// DeepSerializedCopy : Copying through serialization
func (p *Person) DeepSerializedCopy() *Person {
	result := Person{}
	b := bytes.Buffer{}

	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	_ = d.Decode(&result)

	return &result
}

// DeepCopyExample can not be scaled
func DeepCopyExample() {
	john := Person{
		"John",
		&Address{"Johnny Avenue 13", "London", "UK"},
		[]string{"Jim", "Bill"},
	}

	jane := john
	jane.Name = "Jane"
	// INCORRECT:
	//jane.Address.StreetAddress = "321 Baker St"
	//
	//fmt.Println(john.Name, john.Address)
	//fmt.Println(jane.Name, jane. Address)

	jane.Address = &Address{john.Address.StreetAddress, john.Address.City, john.Address.Country}
	jane.Address.StreetAddress = "East West Coast 33"

	fmt.Println(john.Name, john.Address)
	fmt.Println(jane.Name, jane.Address)
}
