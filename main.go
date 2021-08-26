package main

import "fmt"

//
//import "fmt"
//
//// Abstract Factory - is a Creational DP that provides API for creating related family objects.
//
//type Worker interface {
//	getSalary() int
//	setSalary(int)
//	getResponsibilities() string
//}
//
//type Position int
//
//const (
//	DevOpsPosition Position = iota
//	DevPosition
//	AccountantPosition
//)
//
//type Developer struct {
//	Office, Position string
//	salary           int
//}
//
//func (d *Developer) getSalary() int {
//	return d.salary
//}
//
//func (d *Developer) setSalary(salary int) {
//	d.salary = salary
//}
//
//func (d *Developer) getResponsibilities() string {
//	return fmt.Sprintf("I'm watching youtube, write code a bit and receive %d in USD", d.salary)
//}
//
//type DevOps struct {
//	Office, Position string
//	salary           int
//}
//
//func (d *DevOps) getSalary() int {
//	return d.salary
//}
//
//func (d *DevOps) setSalary(salary int) {
//	d.salary = salary
//}
//
//func (d *DevOps) getResponsibilities() string {
//	return fmt.Sprintf("I can autoscale and loadbalance whatever you want. Just pay bills for AWS and give me %d in USD", d.salary)
//}
//
//type Accountant struct {
//	Office, Position string
//	salary           int
//}
//
//func (a *Accountant) getSalary() int {
//	return a.salary
//}
//
//func (a *Accountant) setSalary(salary int) {
//	a.salary = salary
//}
//
//func (a *Accountant) getResponsibilities() string {
//	return fmt.Sprintf("I count big numbers but receive just %d. I need IT courses!", a.salary)
//}
//
//func NewWorker(p Position) (Worker, error) {
//	switch p {
//	case DevOpsPosition:
//		return &DevOps{
//			Office:   "Cozy office with coffee machine",
//			Position: "devops",
//		}, nil
//	case DevPosition:
//		return &Developer{
//			Office:   "Office with smoothie and craft coffee",
//			Position: "developer",
//		}, nil
//	case AccountantPosition:
//		return &Accountant{
//			Office:   "Office with enough light inside",
//			Position: "accountant",
//		}, nil
//	}
//
//	return nil, fmt.Errorf("Wrong position")
//}
//
//func main() {
//	dev1, _ := NewWorker(DevPosition)
//	dev1.setSalary(4500)
//	fmt.Println(dev1.getResponsibilities())
//
//	devOps1, _ := NewWorker(DevOpsPosition)
//	devOps1.setSalary(5000)
//	fmt.Println(devOps1.getResponsibilities())
//
//	acc1, _ := NewWorker(AccountantPosition)
//	acc1.setSalary(500)
//	fmt.Println(acc1.getResponsibilities())
//}
//
// Procedural code
// Procedural code
//type Square struct {
//	size int
//}
//
//type Rectangle struct {
//	height int
//	width int
//}
//
//type Geometry struct {
//}
//
//func (g *Geometry) GetPerimeter(shape interface{}) (int, error) {
//	shapeType := reflect.TypeOf(shape).String()
//	if  shapeType == "main.Square" {
//		square := shape.(Square)
//		return square.size * 4, nil
//	}
//
//	if shapeType == "main.Rectangle" {
//		rectangle := shape.(Rectangle)
//		return rectangle.height * rectangle.width, nil
//	}
//
//	return 0, fmt.Errorf("Not known shape")
//}

// OOP

//type Shape interface {
//	GetPerimeter() int
//}
//type Square struct {
//	size int
//}
//
//func (s *Square) GetPerimeter() int {
//	return s.size * 4
//}
//
//type Rectangle struct {
//	width  int
//	height int
//}
//
//func (r *Rectangle) GetPerimeter() int {
//	return r.height * r.width
//}

//func main() {
//square := Square{10}
//rectangle := Rectangle{10,20}
//geometry := Geometry{}
//squareP, err := geometry.GetPerimeter(square)
//
//if err == nil {
//	fmt.Println(squareP)
//}
////squareR, err := geometry.GetPerimeter(rectangle)
////if err == nil {
////	fmt.Println(squareR)
////}
//
//sq := Square{5}
//rect := Rectangle{10, 15}
//fmt.Println(sq.GetPerimeter())
//fmt.Println(rect.GetPerimeter())
//}

type Human struct {
	Name string
}

type Animal struct {
	Name int
}

type Monster struct {
	Human
	Animal
}

func (m *Monster) PrintName() {
	fmt.Println(m.Human.Name)
}

func main() {
	m := Monster{
		Human:  Human{"Jenia"},
		Animal: Animal{33},
	}

	m.PrintName()
}
