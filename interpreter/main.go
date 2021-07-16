package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/GO_design_patterns/interpreter/examples"
)

// examples:
// Programming language compilers, interpreters and IDEs
// HTML, XML and similar
// Numeric expressions (3+4/5)
// reg.exps
// Interpreter - component that process structured data. Does so by
// turning it into separated lexical tokens (lexing) and
// then interpreting sequences of said tokens (parsing)
// Summary:
// - acts in 2 stages:
// 1.turns smth into set of tokens
// 2. parsing tokens into meaningful constructs
// parsed data can later be traversed using Visitor pattern

func main() {
	input := "(13+4)-(12+1)"
	tokens := interpreter.Lex(input)
	fmt.Println(tokens)

	parsed := interpreter.Parse(tokens)
	fmt.Printf("%s = %d", input, parsed.Value())
}
