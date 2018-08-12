package timesheet

import (
	"fmt"
	"testing"
)

func TestLexer_run(t *testing.T) {
	for _, c := range []struct {
		i     int // which part to assert
		start lexFn
		txt   string
		tok   Token
		val   string
	}{
		{3, lexMonth, "April  \n---\n1", Number, "1"},
		{2, lexMonth, "April  \n  1", Error, "invalid Separator"},
		{1, lexDate, " 4", Error, "invalid Number"},
		{1, lexDate, "4", Number, "4"},
		{2, lexMonth, "August\n3", Error, "invalid Separator"},
		{1, lexWeek, "jkl", Error, "invalid Number"},
		{1, lexWeek, "26", Number, "26"},
		{1, lexYear, "2018", Number, "2018"},
		{1, lexYear, "not a year", Error, "invalid Number"},
		{1, lexMonth, "August", Month, "August"},
		{1, lexMonth, "not a month", Error, "invalid month"},
		{2, lexMonth, "August something more", Error, "expect newline"},
		{1, lexMonth, "Augusty", Error, "invalid month"},
		{1, lexMonth, "august", Error, "invalid month"},
		{1, lexMonth, " August", Error, "invalid month"},
	} {
		l := NewLexer(c.txt)
		out := make(chan Part, 2) // buffer it so we can ignore following parts
		go l.run(c.start, l.scanner, out)
		var part Part
		for i := 0; i < c.i; i++ {
			part = <-out
		}
		assert(t, fmt.Sprintf("%q", c.txt),
			equals("", c.tok, part.Tok),
			equals("", c.val, part.Val),
		)
	}
}

func TestScanPart(t *testing.T) {
	cases := []struct {
		msg, txt string
		tok      Token
	}{
		{"", "1234", Number},
		{"", "as1234", Error},
	}
	for _, c := range cases {
		s := NewScanner(c.txt)
		got := ScanPart(s, Number)
		assert(t, c.msg,
			equals("Tok", c.tok, got.Tok),
		)
	}

}
