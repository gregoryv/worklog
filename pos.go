package timesheet

import (
	"fmt"
)

type Pos struct {
	line, column int
}

func NewPos() *Pos {
	return &Pos{}
}

func (p *Pos) String() string {
	return fmt.Sprintf("%v,%v", p.line, p.column)
}
