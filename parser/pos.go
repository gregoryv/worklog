// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.
package parser

import (
	"fmt"
)

type Position struct {
	line, col int
}

func NewPosition() *Position {
	return &Position{line: 1, col: 1}
}

func (a *Position) Equals(b Position) bool {
	return a.line == b.line &&
		a.col == b.col
}

func (pos *Position) Val() (line, col int) {
	return pos.line, pos.col
}

func (pos Position) String() string {
	return fmt.Sprintf("%v,%v", pos.line, pos.col)
}
