package timesheet

import (
	"testing"
)

func Test_lex(t *testing.T) {
	for _, c := range []struct {
		txt string
		fn  lexFn
		tok Token
		val string
	}{
		{"2018", lexYear, Number, "2018"},
		{"August", lexMonth, Month, "August"},
	} {
		s := NewScanner(c.txt)
		out := make(chan Part)
		go c.fn(s, out)
		part := <-out
		assert(t, "",
			equals("", c.tok, part.Tok),
			equals("", c.val, part.Val),
		)
		close(out)
	}
}
