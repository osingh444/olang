package lexer

import (
	"testing"
)

func TestLexer(t *testing.T) {

	lexer, err := CreateLexer("../test/simple.ol")
	if err != nil {
		panic(err)
	}

	lexer.Lex()
	lexer.PrintTokens()

	lexer2, err := CreateLexer("../test/operators.ol")
	if err != nil {
		panic(err)
	}

	lexer2.Lex()
	lexer2.PrintTokens()
}
