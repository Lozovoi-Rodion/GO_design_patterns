package bridge

import "fmt"

type Renderer interface {
	RenderCircle(radius float32)
	RenderTriangle(a, b, c float32)
}

type VectorRenderer struct {
	//
}

func (v *VectorRenderer) RenderTriangle(a, b, c float32) {
	fmt.Println("Drawing a triangle with sizes ", a, b, c)
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius ", radius)
}

type RasterRenderer struct {
	Dpi int
}

func (r *RasterRenderer) RenderTriangle(a, b, c float32) {
	fmt.Println("Drawing pixels for triangle with sizes", a, b, c)
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius ", radius)
}

type Circle struct {
	renderer Renderer
	radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer, radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

type Triangle struct {
	renderer            Renderer
	sideA, sideB, sideC float32
}

func NewTriangle(renderer Renderer, a,b,c float32) *Triangle {
	return &Triangle{renderer, a,b,c}
}

func (t *Triangle) Draw() {
	t.renderer.RenderTriangle(t.sideA, t.sideB, t.sideC)
}

func (t *Triangle) Resize(factor float32) {
	t.sideA *= factor
	t.sideB *= factor
	t.sideC *= factor
}