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

func TestPart_String(t *testing.T) {
	for _, c := range []struct {
		msg  string
		part Part
		exp  string
	}{
		{"", Part{Tok: Unknown, Val: "jib"}, `Unknown[0,0]: "jib"`},
		{"", Part{Tok: Number, Val: "1"}, `Number[0,0]: "1"`},
		{"Undefined, run 'go generate'", Part{Tok: Token(-1), Val: ""},
			`Token(-1)[0,0]: ""`},
	} {
		got := c.part.String()
		assert(t, c.msg,
			equals("String() return", c.exp, got),
		)
	}
}

func TestNewPart(t *testing.T) {
	part := NewPart()
	if part == nil {
		t.Fail()
	}
}
