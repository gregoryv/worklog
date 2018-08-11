package timesheet

import (
	"fmt"
)

type Pos struct {
	line, column int
}

func NewPos() *Pos {
	return &Pos{line: 1, column: 1}
}

func (p *Pos) Next() (line, column int) {
	p.column++
	return p.line, p.column
}

func (p *Pos) String() string {
	return fmt.Sprintf("%v,%v", p.line, p.column)
}
