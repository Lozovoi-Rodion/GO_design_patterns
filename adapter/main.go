package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/adapter/examples"
)

func main()  {
	rc := adapter.NewRactangle(6,4)
	a := adapter.VectorToRaster(rc)
	fmt.Println(adapter.DrawPoints(a))
}
