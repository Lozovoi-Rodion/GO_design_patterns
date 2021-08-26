package visitor

import (
	"fmt"
	"strings"
)

// violates SRP as it intrudes in all hierarchy

type Expression interface {
	Print(sb *strings.Builder)
}

type DoubleExpression struct {
	Value float64
}

func (d *DoubleExpression) Print(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%g", d.Value))
}

type AdditionExpression struct {
	Left, Right Expression

}

func (a *AdditionExpression) Print(sb *strings.Builder) {
	sb.WriteRune('(')
	a.Left.Print(sb)
	sb.WriteRune('+')
	a.Right.Print(sb)
	sb.WriteRune(')')
}
