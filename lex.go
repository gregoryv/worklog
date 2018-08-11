package timesheet

import (
	"fmt"
)

type Lexer struct {
	name    string // eg. named file
	scanner *Scanner
	out     chan Part
}

func (l *Lexer) Run() chan Part {
	go l.run(l.scanner, l.out)
	return l.out
}

func NewLexer(name, txt string) *Lexer {
	return &Lexer{
		name:    name,
		scanner: NewScanner(txt),
		out:     make(chan Part),
	}
}

type Part struct {
	Tok Token
	Val string
	Pos Position
}

func (p *Part) String() string {
	return fmt.Sprintf("%s[%s]: %q", p.Tok, p.Pos.String(), p.Val)
}

func NewPart() *Part {
	return &Part{}
}
