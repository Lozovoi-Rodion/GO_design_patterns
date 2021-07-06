package prototype

import (
	"bytes"
	"encoding/gob"
)

type Addr struct {
	Suite        int
	Street, City string
}

type Employee struct {
	Name   string
	Office Addr
}

func (p Employee) DeepCopy() *Employee {
	// note: no error handling below
	result := Employee{}
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	_ = d.Decode(&result)

	return &result
}

var mainOffice = Employee{
	"", Addr{0, "77 West Dr", "London"},
}

var auxOffice = Employee{
	"", Addr{0, "66 East", "London"},
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite

	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}
