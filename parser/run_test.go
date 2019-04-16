// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package parser

import (
	"testing"

	"github.com/gregoryv/go-timesheet/token"
)

func Test_ok_lines(t *testing.T) {
	for _, line := range []string{
		"52 24 Mon   Christmas",
		" 1  1 Tue 8",
		"    1 Tue 8 (+1 flex)",
		"    1 Tue 8 (+1 flex) comment (0:30 vacation)",
	} {
		if tok := parse(line, lexWeek); tok == token.Error {
			t.Errorf("%s failed %v", line, tok)
		}
	}
}

func Test_badly_formatted_lines(t *testing.T) {
	for _, line := range []string{
		"Mon   Christmas",
		"tis",
		"\n",
	} {
		if tok := parse(line, lexWeek); tok != token.Error {
			t.Errorf("%s expected to fail", line)
		}
	}
}

func parse(line string, start lexFn) (tok token.Token) {
	lex := NewLexer(line)
	out := lex.C
	go lex.run(start, lex.scanner, out)
	for {
		p, more := <-out
		if p.Tok == token.Error {
			tok = p.Tok
		}
		if !more {
			break
		}
	}
	return
}
