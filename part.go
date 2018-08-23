package timesheet

import (
	"fmt"
)

type Part struct {
	Tok Token
	Val string
	Pos Position
}

func (p *Part) Defined() bool {
	return p.Tok != Undefined
}

func (a Part) Equals(b Part) bool {
	return a.Tok == b.Tok &&
		a.Val == b.Val &&
		a.Pos.Equals(b.Pos)
}

func (p *Part) Errorf(format string, args ...interface{}) error {
	p.Val = fmt.Sprintf(format, args...)
	p.Tok = Error
	return fmt.Errorf(p.Val)
}

func (p Part) String() string {
	return fmt.Sprintf("%s[%s]: %q", p.Tok, p.Pos.String(), p.Val)
}

func NewPart() Part {
	return Part{}
}

func (tok Token) Is(val string, optional ...Position) Part {
	if len(optional) > 0 {
		return Part{tok, val, optional[0]}
	}
	return Part{tok, val, Position{1, 1}}
}
