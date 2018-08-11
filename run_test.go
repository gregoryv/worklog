package timesheet

import (
	"testing"
)

func Test_lex(t *testing.T) {
	for _, c := range []struct {
		txt string
		fn  lexFn
		exp string
	}{
		{"2018", lexYear, `Number[1,1]: "2018"`},
		{"August", lexMonth, `Month[1,1]: "August"`},
	} {
		lex(c.fn, c.txt, c.exp, t)
	}
}

func lex(fn lexFn, txt, exp string, t *testing.T) {
	t.Helper()
	s := NewScanner(txt)
	out := make(chan Part)
	go fn(s, out)
	part := <-out
	assert(t, "",
		equals("", exp, part.String()),
	)
	close(out)
}
