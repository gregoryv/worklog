// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package parser

import "github.com/gregoryv/go-timesheet/token"

type Lexer struct {
	scanner *Scanner
	C       chan Part
}

func NewLexer(txt string) *Lexer {
	return &Lexer{
		scanner: NewScanner(txt),
		C:       make(chan Part),
	}
}

func (l *Lexer) Run() chan Part {
	start := lexYear
	go l.run(start, l.scanner, l.C)
	return l.C
}

type lexFn func(s *Scanner) (p Part, next lexFn)

func (l *Lexer) run(start lexFn, s *Scanner, C chan Part) {
	p, next := start(s)
	for {
		if p.Tok != token.Undefined {
			C <- p
		}
		if next == nil {
			break
		}
		p, next = next(s)
	}
	close(C)
}
