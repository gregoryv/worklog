package timesheet

import (
	. "github.com/gregoryv/qual"
	"testing"
)

func TestLexer_Run(t *testing.T) {
	l := NewLexer("2018")
	out := l.Run()
	cases := []struct {
		part Part
		exp  string
	}{
		{<-out, "Number[1,1]: \"2018\""},
	}
	for _, c := range cases {
		got := c.part.String()
		Assert(t, Vars{c.exp, got}, c.exp == got)
	}
}

func TestNewLexer(t *testing.T) {
	if l := NewLexer(""); l == nil {
		t.Fail()
	}
}
