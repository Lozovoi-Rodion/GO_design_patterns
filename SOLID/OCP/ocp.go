package OCP

// open for extension, closed for modification
// Specification

type Color int

const (
	Red Color = iota
	Green
	Blue
)

type Size int

const (
	Small Size = iota
	Medium
	Large
)

type Product struct {
	Name  string
	Color Color
	Size  Size
}

type Filter struct {}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.Color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

/*Not correct to extend variety of filters, Specification pattern has to be used*/
func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.Size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	Color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.Color == c.Color
}

type SizeSpecification struct {
	Size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.Size == s.Size
}

type AndSpecification struct {
	First, Second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool{
	return a.First.IsSatisfied(p) && a.Second.IsSatisfied(p)
}

type OrSpecification struct {
	First, Second Specification
}

func (or OrSpecification) IsSatisfied(p *Product) bool{
	return or.First.IsSatisfied(p) || or.Second.IsSatisfied(p)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}
