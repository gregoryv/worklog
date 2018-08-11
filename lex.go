package timesheet

import (
	"fmt"
)

type Lexer struct {
	scanner *Scanner
}

func NewLexer(txt string) *Lexer {
	return &Lexer{
		scanner: NewScanner(txt),
	}
}

type Part struct {
	tok Token
	val string
	pos Position
}

func (p *Part) String() string {
	return fmt.Sprintf("%s[%s]: %q", p.tok, p.pos.String(), p.val)
}

func NewPart() *Part {
	return &Part{}
}
