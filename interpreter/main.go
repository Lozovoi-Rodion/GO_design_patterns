package main

// examples:
// Programming language compilers, interpreters and IDEs
// HTML, XML and similar
// Numeric expressions (3+4/5)
// reg.exps
// Interpreter - component that process structured data. Does so by
// turning it into separated lexical tokens (lexing) and
// then interpreting sequences of said tokens (parsing)
func main() {
	input := "(13+4)-(12+1)"
	tokens := interpreter.Lex(input)
}
