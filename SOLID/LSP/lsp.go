package LSP

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	Width, Height int
}

//     vvv !! POINTER
func (r *Rectangle) GetWidth() int {
	return r.Width
}

func (r *Rectangle) SetWidth(width int) {
	r.Width = width
}

func (r *Rectangle) GetHeight() int {
	return r.Height
}

func (r *Rectangle) SetHeight(height int) {
	r.Height = height
}

// modified LSP
// If a function takes an interface and
// works with a type T that implements this
// interface, any structure that aggregates T
// should also be usable in that function.
type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.Width = size
	sq.Height = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.Width = width
	s.Height = width
}

func (s *Square) SetHeight(height int) {
	s.Width = height
	s.Height = height
}

type Square2 struct {
	Size int
}

func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.Size, s.Size}
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected an area of ", expectedArea,
		", but got ", actualArea, "\n")
}