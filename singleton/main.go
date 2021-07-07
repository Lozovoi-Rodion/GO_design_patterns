package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/singleton/examples"
)

func main()  {
	db := singleton.GetSingletonDatabase()
	pop := db.GetPopulation("Mumbai")
	fmt.Println("population of Mumbai is ", pop)
}
