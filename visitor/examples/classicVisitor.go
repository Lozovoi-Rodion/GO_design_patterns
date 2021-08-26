package visitor

import (
	"fmt"
	"strings"
)

type ExpVisitor interface {
	VisitDoubleExp(de *DoubleExp)
	VisitAdditionExp(ae *AdditionExp)
}

type Exp interface {
	Accept(ev ExpVisitor)
}

type AdditionExp struct {
	Left, Right Exp
}

func (ae *AdditionExp) Accept(ev ExpVisitor) {
	ev.VisitAdditionExp(ae)
}

type DoubleExp struct {
	Value float64
}

func (de *DoubleExp) Accept(ev ExpVisitor) {
	ev.VisitDoubleExp(de)
}

type ExpressionPrinter struct {
	sb strings.Builder
}

func (ep *ExpressionPrinter) VisitDoubleExp(de *DoubleExp) {
	ep.sb.WriteString(fmt.Sprintf("%g", de.Value))
}

func (ep *ExpressionPrinter) VisitAdditionExp(ae *AdditionExp) {
	ep.sb.WriteRune('(')
	ae.Left.Accept(ep)
	ep.sb.WriteRune('+')
	ae.Right.Accept(ep)
	ep.sb.WriteRune(')')
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{strings.Builder{}}
}

func (ep *ExpressionPrinter) String() string {
	return ep.sb.String()
}
