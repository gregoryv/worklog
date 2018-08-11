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

func (pos *Position) Back() (line, col int) {
	line, col = pos.line, pos.col
	switch {
	case col > 1 && line >= 1:
		col--
	case col == 1 && line > 1:
		if pos.last == -1 {
			panic("You can only Back once over a newline")
		}
		col = pos.last
		pos.last = -1
		line--
	}
	pos.line = line
	pos.col = col
	return
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
