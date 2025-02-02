package tokens

import (
	"errors"
	"fmt"
	"strconv"
)

type TokenType uint8

const (
	None TokenType = iota
	Number
	String
	Char
	Bool
	Ident
	Funcall
	OpenBracket
	CloseBracket
	OpenParen
	CloseParen
	Fun
	If
	Else
	Semicolon
	Assign
	Pipe
	Plus
	Concat
	Hyphen
	Asterisk
	ForwardSlash
	Percent
	Equals
	NotEquals
	GreaterThan
	LesserThan
	GreaterThanOrEqualTo
	LesserThanOrEqualTo
	And
	Or
	BitAnd
	BitOr
	BitXOR
)

func (tt TokenType) Str() string {
	return []string{
		"None",
		"Number",
		"String",
		"Char",
		"Bool",
		"Ident",
		"Funcall",
		"OpenBracket",
		"CloseBracket",
		"OpenParen",
		"CloseParen",
		"Fun",
		"If",
		"Else",
		"Semicolon",
		"Assign",
		"Pipe",
		"Plus",
		"Hyphen",
		"Asterisk",
		"ForwardSlash",
		"Percent",
		"Equals",
		"NotEquals",
		"GreaterThan",
		"LesserThan",
		"GreaterThanOrEqualTo",
		"LesserThanOrEqualTo",
		"And",
		"Or",
		"BitAnd",
		"BitOr",
		"BitXOR",
	}[tt]
}

type Token struct {
	kind      TokenType
	lit       string
	start, ln int
}

func New(kind TokenType, lit string, start, ln int) Token {
	return Token{kind: kind, lit: lit, start: start, ln: ln}
}

func NewLit(kind TokenType, lit string) Token {
	t := Empty()
	t.kind = kind
	t.lit = lit
	return t
}

func Empty() Token {
	return New(None, "", -1, -1)
}

func (t Token) GetKind() TokenType {
	return t.kind
}

func (t Token) GetLit() string {
	return t.lit
}

func (t Token) GetLn() int {
	return t.ln
}

func (t Token) IsKind(kind TokenType) bool {
	return t.kind == kind
}

func (t Token) IsLit(lit string) bool {
	return t.lit == lit
}

func (t Token) Convert() (any, error) {
	switch t.kind {
	case Number:
		{
			f, err := strconv.ParseFloat(t.lit, 32)
			return float32(f), err
		}
	case String:
		return t.lit, nil
	case Bool:
		return t.lit == "True", nil
	}
	panic("invalid kind: " + t.Str())
}

func (t Token) Str() string {
	return fmt.Sprintf("{%s, '%s', %d, %d}", t.kind.Str(), t.lit, t.start, t.ln)
}

func (t Token) Err(format string, a ...any) error {
	return errors.New(fmt.Sprintf("error on line %d, col %d: ", t.ln, t.start) + fmt.Sprintf(format, a...))
}
