package timesheet

import (
	"testing"
)

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
		{"", Part{tok: Unknown, val: "jib"}, `Unknown[0,0]: "jib"`},
		{"", Part{tok: Number, val: "1"}, `Number[0,0]: "1"`},
		{"Undefined, run 'go generate'", Part{tok: Token(-1), val: ""},
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
