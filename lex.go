package timesheet

import (
	"fmt"
)

type Part struct {
	tok Token
	val string
}

func (p *Part) String() string {
	return fmt.Sprintf("%s: %q", p.tok, p.val)
}

func NewPart() *Part {
	return &Part{}
}
