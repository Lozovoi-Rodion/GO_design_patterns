package main

import "github.com/Lozovoi-Rodion/GO_design_patterns/bridge/examples"
// Bridge prevents a 'Cartesian product' complexity explosion
// Example:
// - Common type ThreadScheduler
// - can be preemptive or cooperative
// - can run on Windows or Unix
// Bride is a mechanism that decouples an interface(hierarchy) from an implementation (hierarchy)

// Summary:
// - decouple abstraction from implementation
// - both can exist as hierarchies
// - a stronger form of encapsulation

func main() {
	raster := bridge.RasterRenderer{}
	vector := bridge.VectorRenderer{}

	circleR := bridge.NewCircle(&raster, 5)
	circleR.Draw()
	circleR.Resize(2)
	circleR.Draw()

	circleV := bridge.NewCircle(&vector, 5)
	circleV.Draw()
	circleV.Resize(2)
	circleV.Draw()

	triangleR := bridge.NewTriangle(&raster, 1,2,3.5)
	triangleR.Draw()
	triangleR.Resize(3)
	triangleR.Draw()

	triangleV := bridge.NewTriangle(&vector, 2,4,7)
	triangleV.Draw()
	triangleV.Resize(5)
	triangleV.Draw()
}