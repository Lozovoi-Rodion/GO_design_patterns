package visitor

import (
	"fmt"
	"strings"
)
type Expr interface {
	// nothing here!
}

type DoubleExpr struct {
	Value float64
}

type AdditionExpr struct {
	Left, Right Expr
}

func Print(e Expr, sb *strings.Builder) {
	if de, ok := e.(*DoubleExpr); ok {
		sb.WriteString(fmt.Sprintf("%g", de.Value))
	} else if ae, ok := e.(*AdditionExpr); ok {
		sb.WriteString("(")
		Print(ae.Left, sb)
		sb.WriteString("+")
		Print(ae.Right, sb)
		sb.WriteString(")")
	}

	// breaks OCP
	// will work incorrectly on missing case
}