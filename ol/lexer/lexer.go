package lexer

import (
	"fmt"
	"io/ioutil"
)

type Lexer struct {
	File     string
	Length   int
	Start    int
	Position int
	Tokens   []*Token
}

func CreateLexer(filename string) (*Lexer, error) {

	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var tokens []*Token

	strfile := string(file[:])
	lexer := Lexer{File: strfile, Length: len(strfile), Start: 0, Position: 0, Tokens: tokens}

	return &lexer, nil
}

func (lexer *Lexer) Lex() {

	isEOF := false

	for !isEOF {
		token := lexer.advanceToken()

		if token != nil {
			lexer.addToken(token)
			if token.Type == EOF {
				isEOF = true
			}
		}

		lexer.Start = lexer.Position
	}
}

func (lexer *Lexer) advanceToken() *Token {
	if lexer.isAtEnd() {
		return CreateToken(lexer, EOF)
	}

	switch ch := lexer.advance(); ch {

	case "("[0]:
		return CreateToken(lexer, LPAREN)
	case ")"[0]:
		return CreateToken(lexer, RPAREN)
	case "{"[0]:
		return CreateToken(lexer, LBRACK)
	case "}"[0]:
		return CreateToken(lexer, RBRACK)
	case ";"[0]:
		return CreateToken(lexer, SEMICOLON)
	case "."[0]:
		return CreateToken(lexer, PERIOD)
	case ","[0]:
		return CreateToken(lexer, COMMA)
	case "+"[0]:
		return CreateToken(lexer, PLUS)
	case "-"[0]:
		return CreateToken(lexer, MINUS)
	case "*"[0]:
		return CreateToken(lexer, TIMES)

	//2 character cases
	case "/"[0]:
		if lexer.nextMatches("/") {
			return lexer.comment()
		}
		return CreateToken(lexer, SLASH)
	case "!"[0]:
		if lexer.nextMatches("=") {
			return CreateToken(lexer, NOTEQ)
		}
		return CreateToken(lexer, NOT)
	case "="[0]:
		if lexer.nextMatches("=") {
			return CreateToken(lexer, EQEQ)
		}
		return CreateToken(lexer, EQ)
	case "<"[0]:
		if lexer.nextMatches("=") {
			return CreateToken(lexer, LTEQ)
		}
		return CreateToken(lexer, LT)
	case ">"[0]:
		if lexer.nextMatches("=") {
			return CreateToken(lexer, GTEQ)
		}
		return CreateToken(lexer, GT)
	case "\""[0]:
		return lexer.str()

	//skip whitespace
	case " "[0]:
		return nil
	case "\n"[0]:
		return nil
	case "\t"[0]:
		return nil
	case "\r"[0]:
		return nil

	default:
		if isNumeric(ch) {
			return lexer.number()
		} else if isAlpha(ch) {
			return lexer.identifier(ch)
		}
		break
	}

	return CreateToken(lexer, UNRECOGNIZED)
}

func (lexer *Lexer) addToken(token *Token) {
	lexer.Tokens = append(lexer.Tokens, token)
}

func (lexer *Lexer) number() *Token {
	for isNumeric(lexer.peek(0)) {
		lexer.advance()
	}
	return CreateToken(lexer, NUMBER)
}

func (lexer *Lexer) comment() *Token {
	for !lexer.isAtEnd() && lexer.advance() != "\n"[0] {
	}
	return nil
}

func (lexer *Lexer) str() *Token {
	for lexer.advance() != "\""[0] {
	}
	return CreateToken(lexer, STRING)
}

func (lexer *Lexer) identifier(start byte) *Token {

	isKW, keyword := lexer.isKeyword(start)
	if isKW {
		return CreateToken(lexer, keyword)
	}

	for isNumeric(lexer.peek(0)) || isAlpha(lexer.peek(0)) {
		lexer.advance()
	}
	return CreateToken(lexer, IDENTIFIER)
}

func (lexer *Lexer) isKeyword(start byte) (bool, int) {

	switch start {
	case "t"[0]:
		if lexer.nextMatches("rue") {
			return true, TRUE
		}
	case "f"[0]:
		if lexer.nextMatches("alse") {
			return true, FALSE
		} else if lexer.nextMatches("or") {
			return true, FOR
		} else if lexer.nextMatches("unc") {
			return true, FUNC
		}
	case "a"[0]:
		if lexer.nextMatches("nd") {
			return true, AND
		}
	case "o"[0]:
		if lexer.nextMatches("r") {
			return true, OR
		}
	case "i"[0]:
		if lexer.nextMatches("f") {
			return true, IF
		}
	case "e"[0]:
		if lexer.nextMatches("lse") {
			return true, ELSE
		}
	case "w"[0]:
		if lexer.nextMatches("hile") {
			return true, WHILE
		}
	case "v"[0]:
		if lexer.nextMatches("ar") {
			return true, VAR
		}
	}


	return false, 0
}

func (lexer *Lexer) isAtEnd() bool {
	return lexer.Position == lexer.Length - 1
}

func isAlpha(ch byte) bool {
	return (ch >= 65 && ch <= 90) || (ch >= 97 && ch <= 122)
}

func isNumeric(ch byte) bool {
	return ch <= 57 && ch >= 48
}

func (lexer *Lexer) advance() byte {
	if lexer.isAtEnd() {
		return 0
	}

	lexer.Position += 1
	return lexer.File[lexer.Position - 1]
}

func (lexer *Lexer) peek(i int) byte {
	if lexer.isAtEnd() {
		return 0
	}

	return lexer.File[lexer.Position + i]
}

func (lexer *Lexer) nextMatches(chrs string) bool {

	for i, ch := range chrs {
		if ch != rune(lexer.peek(i)) {
			return false
		}
	}

	for _, _ = range chrs {
		lexer.advance()
	}

	return true
}

func (lexer *Lexer) PrintTokens() {
	for _, token := range lexer.Tokens {
		fmt.Println(token.Value, token.Length, token.Type)
	}
}
