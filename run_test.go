package timesheet

import (
	"testing"
)

func Test_oklines(t *testing.T) {
	// not the first two
	lines := []string{
		"52 24 Mon   Christmas",
	}
	for _, line := range lines {
		lex := NewLexer(line)
		out := lex.C
		go lex.run(lexWeek, lex.scanner, out)
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
}

func Test_lexMinutes_inside_tag(t *testing.T) {
	input := "10"
	scanner := NewLexer(input).scanner
	scanner.inTag = true
	got, _ := lexMinutes(scanner)
	exp := Minutes.Is("10")
	if got != exp {
		t.Errorf("input %q, scanned as %q, expected %q", input, got, exp)
	}
}

func TestLexer_run(t *testing.T) {
	for _, c := range []struct {
		start lexFn
		input string
		exp   Part
	}{
		{lexTag, "k", Error.Is("missing RightParenthesis", Position{1, 2})},
		{lexTag, "word (", Error.Is("missing RightParenthesis", Position{1, 7})},
		{lexTag, " vacation hej)", Tag.Is("vacation hej")},

		{lexMinutes, "0", Minutes.Is("0")},
		{lexMinutes, "-10", Error.Is("invalid Minutes")},
		{lexMinutes, "aa", Error.Is("invalid Minutes")},

		{lexColon, "jkjk", Undefined.Is("")},
		{lexColon, ":", Colon.Is(":")},

		{lexHours, "-", Error.Is("invalid Hours")},
		{lexHours, "", Error.Is("invalid Hours")},
		{lexHours, "2", Hours.Is("2")},

		{lexOperator, "+", Operator.Is("+")},
		{lexOperator, "-", Operator.Is("-")},
		{lexOperator, "k", Error.Is("invalid Operator")},

		{lexLeftParen, "kj", Error.Is("invalid LeftParenthesis")},
		{lexLeftParen, "(", LeftParenthesis.Is("(")},
		{lexNote, "(8 working)", Undefined.Is("")},
		{lexNote, "not working\n", Note.Is("not working")},
		{lexNote, "  not working", Note.Is("  not working")},

		{lexDay, "Mo", Error.Is("invalid Day")},
		{lexDay, "mon", Error.Is("invalid Day")},
		{lexDay, "Wen", Error.Is("invalid Day")},
		{lexDay, "Mon", Day.Is("Mon")},
		{lexDay, "Wed", Day.Is("Wed")},

		{lexDate, " 4", Error.Is("invalid Date")},
		{lexDate, "4", Date.Is("4")},

		{lexWeek, "26   1", Week.Is("26")},
		{lexWeek, "     2", Undefined.Is("", Position{0, 0})},
		{lexWeek, "jkl", Error.Is("invalid Week")},
		{lexWeek, "26", Week.Is("26")},

		{lexYear, "2018", Year.Is("2018")},
		{lexYear, "not a year", Error.Is("invalid Year")},

		{lexSep, "-----", Separator.Is("-----")},

		{lexMonth, "April  \n---\n11", Month.Is("April")},
		{lexMonth, "August\n", Month.Is("August")},
		{lexMonth, "not a month", Error.Is("invalid Month")},
		{lexMonth, "August something more",
			Error.Is("expect newline", Position{1, 7}),
		},
		{lexMonth, "Augusty", Error.Is("invalid Month")},
		{lexMonth, "august", Error.Is("invalid Month")},
		{lexMonth, " August", Error.Is("invalid Month")},

		{lexRightParen, "not", Error.Is("invalid RightParenthesis")},
	} {
		input, exp := c.input, c.exp
		got, _ := c.start(NewLexer(c.input).scanner)
		if got != exp {
			t.Errorf("%q, scanned as\n%q, expected\n%q", input, got, c.exp)
		}
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
		exp      Token
	}{
		{"", "1234", Year},
		{"", "as1234", Error},
	}
	for _, c := range cases {
		s := NewScanner(c.txt)
		got := ScanPart(s, Year)
		exp := c.exp
		if got.Tok != exp {
			t.Errorf("%q, expected %q", got, exp)
		}
	}

}
