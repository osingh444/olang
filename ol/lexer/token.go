package lexer

type Token struct {
	Value  string
	Type   int
	Length int
}

//declare the different token types
const (
	NUMBER = iota
	IDENTIFIER = iota
	STRING = iota

	LPAREN = iota
	RPAREN = iota
	LBRACK = iota
	RBRACK = iota
	COMMA = iota
	PERIOD = iota
	SEMICOLON = iota
	PLUS = iota
	MINUS = iota
	TIMES = iota
	SLASH = iota

	GT = iota
	GTEQ = iota
	LT = iota
	LTEQ = iota
	EQ = iota
	EQEQ = iota
	NOT = iota
	NOTEQ = iota

	TRUE = iota
	FALSE = iota

	AND = iota
	OR = iota
	IF = iota
	ELSE = iota
	FUNC = iota

	FOR = iota
	WHILE = iota

	VAR = iota

	EOF = iota
	UNRECOGNIZED = iota
)

func CreateToken(lexer *Lexer, typ int) (*Token) {

	token := Token{Value: lexer.File[lexer.Start: lexer.Position], Type: typ}

	token.Length = lexer.Position - lexer.Start

	return &token
}
