package timesheet

import (
	"fmt"
)

type Position struct {
	line, col int
}

func NewPosition() *Position {
	return &Position{line: 1, col: 1}
}

func (pos *Position) Back() (line, col int) {
	if pos.col > 1 {
		pos.col--
	}
	if pos.line > 1 {
		pos.line--
	}
	return pos.line, pos.col
}

func (pos *Position) NextLine() (line, col int) {
	pos.line++
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
