package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/singleton/examples"
)

//Summary:
// - lazy one-time initialization using sync.Once
// - Adhere to DIP: depend on interfaces, not concrete types

func main() {
	db := singleton.GetSingletonDatabase()
	pop := db.GetPopulation("Mumbai")
	fmt.Println("population of Mumbai is ", pop)

	cities := []string{"Seoul", "Mexico City"}
	//tp := singleton.GetTotalPopulation(cities) - inappropriate way of testing as we rely on Low level module. We should rely on interface
	tp := singleton.GetTotalPopulationEx(singleton.GetSingletonDatabase(), cities)
	ok := tp == (17500000 + 17400000) // testing on live data
	fmt.Println(ok)

	names := []string{"alpha", "gamma"} // expect 4
	tp = singleton.GetTotalPopulationEx(&DummyDatabase{}, names)
	ok = tp == 4
	fmt.Println(ok)
}
