package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/visitor/examples"
	"strings"
)

// Visitor - dp where a component (visitor) is allowed to traverse the entire hierarchy
// of types. Implemented by propagating a single Accept() method throughout tje entire hierarchy
// Summary:
// - propagate an Accept(v *Visitor) through entire hierarchy
// - create a visitor with a bunch of methods (VisitA(f A), visitB(f B) for all elements in hierarchy
// - each Accept simply calls Visitor.VisitX(v X) method
func main() {
	// 1 + (2+3)
	e := &visitor.AdditionExpression{
		Left: &visitor.DoubleExpression{1},
		Right: &visitor.AdditionExpression{
			Left:  &visitor.AdditionExpression{
				Left:  &visitor.DoubleExpression{22},
				Right: &visitor.DoubleExpression{23},
			},
			Right: &visitor.DoubleExpression{3},
		},
	}
	sb := strings.Builder{}
	// string builder is actual visitor
	e.Print(&sb)
	fmt.Println(sb.String())

	sb.Reset()
	e2 := &visitor.AdditionExpr{
		Left: &visitor.DoubleExpr{1},
		Right: &visitor.AdditionExpr{
			Left:  &visitor.DoubleExpr{2},
			Right: &visitor.AdditionExpr{
				Left:  &visitor.DoubleExpr{31},
				Right: &visitor.DoubleExpr{32},
			},
		},
	}
	visitor.Print(e2, &sb)
	fmt.Println(sb.String())

	e3 := &visitor.AdditionExp{
		Left: &visitor.DoubleExp{1},
		Right: &visitor.AdditionExp{
			Left: &visitor.DoubleExp{5},
			Right: &visitor.DoubleExp{4},
		},
	}
	ep := visitor.NewExpressionPrinter()
	e3.Accept(ep)
	fmt.Println(ep.String())
}
