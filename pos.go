package timesheet

import (
	"fmt"
)

type Position struct {
	line, col int
	last      int // last column when using NextLine so we can Back once
}

func NewPosition() *Position {
	return &Position{line: 1, col: 1, last: 1}
}

func (pos *Position) Val() (line, col int) {
	return pos.line, pos.col
}

func (pos Position) String() string {
	return fmt.Sprintf("%v,%v", pos.line, pos.col)
}
