package parser

import (
	"testing"
)

func Test_ok_lines(t *testing.T) {
	for _, line := range []string{
		"52 24 Mon   Christmas",
		" 1  1 Tue 8",
		"    1 Tue 8 (+1 flex)",
		"    1 Tue 8 (+1 flex) comment (0:30 vacation)",
	} {
		if tok := parse(line, lexWeek); tok == Error {
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
		if tok := parse(line, lexWeek); tok != Error {
			t.Errorf("%s expected to fail", line)
		}
	}
}

func parse(line string, start lexFn) (tok Token) {
	lex := NewLexer(line)
	out := lex.C
	go lex.run(start, lex.scanner, out)
	for {
		p, more := <-out
		if p.Tok == Error {
			tok = p.Tok
		}
		if !more {
			break
		}
	}
	return
}
