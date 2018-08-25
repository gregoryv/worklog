package timesheet

import (
	. "github.com/gregoryv/qual"
	"testing"
)

func TestLexer_run(t *testing.T) {
	for _, c := range []struct {
		start lexFn
		input string
		exp   Part
	}{
		{lexTag, "k", Error.Is("missing RightParenthesis", Position{1, 2})},
		{lexTag, " vacation hej)", Tag.Is("vacation hej")},

		{lexMinutes, "0", Minutes.Is("0")},
		{lexMinutes, "-10", Error.Is("invalid Minutes")},
		{lexMinutes, "aa", Error.Is("invalid Minutes")},

		{lexColon, "jkjk", Undefined.Is("")},
		{lexColon, ":", Colon.Is(":")},

		{lexHours, "-", Error.Is("invalid Hours")},
		{lexHours, "", Error.Is("invalid Hours")},
		{lexHours, "2", Hours.Is("2")},

		{lexOperator, "8", Undefined.Is("")},
		{lexOperator, " ", Error.Is("invalid Operator")},
		{lexOperator, "+", Operator.Is("+")},
		{lexOperator, "-", Operator.Is("-")},

		{lexLeftParen, "kj", Error.Is("invalid LeftParenthesis")},
		{lexLeftParen, "(", LeftParenthesis.Is("(")},
		{lexNote, "(8 working)", Undefined.Is("")},
		{lexNote, "not working\n", Note.Is("not working\n")},
		{lexNote, "not working", Note.Is("not working")},
		{lexReported, "8 what a day", Hours.Is("8")},
		{lexReported, "6 (", Hours.Is("6")},
		{lexReported, "6\n", Hours.Is("6")},
		{lexReported, "\n", Undefined.Is("")},
		{lexReported, " ", Error.Is("invalid Hours")},

		{lexDay, "Mo", Error.Is("invalid Day")},
		{lexDay, "mon", Error.Is("invalid Day")},
		{lexDay, "Mon", Day.Is("Mon")},

		{lexDate, " 4", Error.Is("invalid Date")},
		{lexDate, "4", Date.Is("4")},

		{lexWeek, "26   1", Number.Is("26")},
		{lexWeek, "     2", Undefined.Is("", Position{0, 0})},
		{lexWeek, "jkl", Error.Is("invalid Number")},
		{lexWeek, "26", Number.Is("26")},

		{lexYear, "2018", Year.Is("2018")},
		{lexYear, "not a year", Error.Is("invalid Year")},

		{lexSep, "-----", Separator.Is("-----")},

		{lexMonth, "April  \n---\n11", Month.Is("April")},
		{lexMonth, "August\n", Month.Is("August")},
		{lexMonth, "not a month", Error.Is("invalid month")},
		{lexMonth, "August something more",
			Error.Is("expect newline", Position{1, 7}),
		},
		{lexMonth, "Augusty", Error.Is("invalid month")},
		{lexMonth, "august", Error.Is("invalid month")},
		{lexMonth, " August", Error.Is("invalid month")},
	} {
		input, exp := c.input, c.exp
		got, _ := c.start(NewLexer(c.input).scanner)
		Assert(t, Vars{input, exp, got},
			got.Equals(exp),
		)
	}
}

func skipParts(i int, out chan Part) (p Part) {
	for j := 0; j < i; j++ {
		p = <-out
	}
	return
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
		Assert(t, Vars{c, got}, c.tok == got.Tok)
	}

}
