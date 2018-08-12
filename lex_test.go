package timesheet

import (
	"testing"
)

func TestLexer_Run(t *testing.T) {
	l := NewLexer("", "2018")
	out := l.Run()
	cases := []struct {
		part Part
		exp  string
	}{
		{<-out, "Number[1,1]: \"2018\""},
	}
	for _, c := range cases {
		got := c.part.String()
		assert(t, "",
			equals("Part.String()", c.exp, got),
		)
	}
}

func TestNewLexer(t *testing.T) {
	if l := NewLexer("lex_test.go", ""); l == nil {
		t.Fail()
	}
}
