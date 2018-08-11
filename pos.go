package timesheet

import (
	"fmt"
)

type Pos struct {
	line, col int
}

func NewPos() *Pos {
	return &Pos{line: 1, col: 1}
}

func (p *Pos) NextLine() (line, col int) {
	p.line++
	p.col = 1
	return p.line, p.col
}

func (p *Pos) Next() (line, col int) {
	p.col++
	return p.line, p.col
}

func (p *Pos) String() string {
	return fmt.Sprintf("%v,%v", p.line, p.col)
}
