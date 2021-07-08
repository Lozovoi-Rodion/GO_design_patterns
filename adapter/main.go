package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/adapter/examples"
)

// Summary:
// - determine the API you have and the API you need
// - create a component which aggregates (has a pointer to, ...) the adaptee
// - intermediate representations can pile up: use caching and other optimizations

func main()  {
	rc := adapter.NewRectangle(25,5)
	a := adapter.VectorToRaster(rc)
	_ = adapter.VectorToRaster(rc)
	fmt.Println(adapter.DrawPoints(a))
}
