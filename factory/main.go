package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/factory/examples"
)

//Summary:
// - a factory function (a.k.a. constructor) is a helper for making struct instances
// - a factory is any entity that can take care of object creation
// - can be a f-n or dedicated struct

func main() {
	p1 := factory.NewPerson1("John", 22)
	p1.EyeCount = 4
	fmt.Println(p1)

	p2 := factory.NewPerson2("Jimmy", 20)
	p2.SayHello()
	p3 := factory.NewPerson2("Derick", 101)
	p3.SayHello()

	developerFactory := factory.NewEmployeeFactory("developer", 60000)
	managerFactory := factory.NewEmployeeFactory("manager", 80000)
	devOpsFactory := factory.NewEmployeeFactory2("devOps", 70000)

	dev := developerFactory("Andy")
	manager := managerFactory("Jane")
	sre := devOpsFactory.Create("Boris")

	fmt.Println(dev)
	fmt.Println(manager)
	fmt.Println(sre)

	developer := factory.NewWorker(factory.Developer)
	developer.Name = "Joshua"
	developer.Company = "Apple"

	fmt.Println(developer)
}
