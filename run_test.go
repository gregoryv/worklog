package timesheet

import (
	"fmt"
	"testing"
)

func Test_lex(t *testing.T) {
	for _, c := range []struct {
		i   int // which part to read
		fn  lexFn
		txt string
		tok Token
		val string
	}{
		{1, lexWeek, "jkl", Error, "invalid week"},
		{1, lexWeek, "26", Number, "26"},
		{1, lexYear, "2018", Number, "2018"},
		{1, lexYear, "not a year", Error, "invalid year"},
		{1, lexMonth, "August", Month, "August"},
		{1, lexMonth, "not a month", Error, "invalid month"},
		{2, lexMonth, "August something more", Error, "line should end"},
		{1, lexMonth, "Augusty", Error, "invalid month"},
		{1, lexMonth, "august", Error, "invalid month"},
		{1, lexMonth, " August", Error, "invalid month"},
	} {
		s := NewScanner(c.txt)
		out := make(chan Part, 3) // buffer it so we can ignore following parts
		go c.fn(s, out)
		var part Part
		for i := 0; i < c.i; i++ {
			part = <-out
		}
		assert(t, fmt.Sprintf("%q", c.txt),
			equals("", c.tok, part.Tok),
			equals("", c.val, part.Val),
		)
		close(out)
	}
}
