package timesheet

import (
	"fmt"
)

type Position struct {
	line, col int
	last      int // last column when using NextLine so we can Back once
}

func NewPosition() *Position {
	return &Position{line: 1, col: 1}
}

func (pos *Position) Val() (line, col int) {
	return pos.line, pos.col
}

func (pos *Position) Back() (line, col int) {
	if pos.col == 1 && pos.line > 1 {
		if pos.last == -1 {
			panic("You can only Back once over a newline")
		}
		pos.col = pos.last
		pos.last = -1
	} else if pos.col > 1 {
		pos.col--
	}
	if pos.line > 1 {
		pos.line--
	}
	return pos.line, pos.col
}

func (pos *Position) NextLine() (line, col int) {
	pos.line++
	pos.last = pos.col
	pos.col = 1
	return pos.line, pos.col
}

func (pos *Position) Next() (line, col int) {
	pos.col++
	return pos.line, pos.col
}

func (pos *Position) String() string {
	return fmt.Sprintf("%v,%v", pos.line, pos.col)
}
