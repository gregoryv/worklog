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
		okLine(t, line, lexWeek)
	}
}

func Test_badly_formatted_lines(t *testing.T) {
	for _, line := range []string{
		"Mon   Christmas",
		"tis",
		"\n",
	} {
		badLine(t, line, lexWeek)
	}
}

func badLine(t *testing.T, line string, start lexFn) {
	t.Helper()
	lex := NewLexer(line)
	out := lex.C
	go lex.run(start, lex.scanner, out)
	var gotErr bool
	for {
		p, more := <-out
		if p.Tok == Error {
			gotErr = true
		}
		if !more {
			break
		}
	}
	if !gotErr {
		t.Errorf("%q expected to fail", line)
	}
}

func okLine(t *testing.T, line string, start lexFn) {
	t.Helper()
	lex := NewLexer(line)
	out := lex.C
	go lex.run(start, lex.scanner, out)
	for {
		p, more := <-out
		if p.Tok == Error {
			t.Errorf("%q, got %q", line, p.Val)
		}
		if !more {
			break
		}
	}
}
