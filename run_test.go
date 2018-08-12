package timesheet

import (
	"testing"
)

func Test_lex(t *testing.T) {
	out := make(chan Part)
	for _, c := range []struct {
		txt string
		fn  lexFn
		tok Token
		val string
	}{
		{"2018", lexYear, Number, "2018"},
		{"not a year", lexYear, Error, "invalid year"},
		{"August", lexMonth, Month, "August"},
		{"not a month", lexMonth, Error, "invalid month"},
		{"Augusty", lexMonth, Error, "invalid month"},
	} {
		s := NewScanner(c.txt)
		go c.fn(s, out)
		part := <-out
		assert(t, "",
			equals("", c.tok, part.Tok),
			equals("", c.val, part.Val),
		)
	}
	close(out)
}
