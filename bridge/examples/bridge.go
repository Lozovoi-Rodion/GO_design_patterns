package bridge

import "fmt"

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRender struct {
	//
}

func (v *VectorRender) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius ", radius)
}

type RasterRender struct {
	Dpi int
}

func (r *RasterRender) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius ", radius)
}

type Circle struct {
	renderer Renderer
	radius float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer, radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32)  {
	c.radius *= factor
}
