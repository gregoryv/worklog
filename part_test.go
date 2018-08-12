package timesheet

import (
	"testing"
)

func TestPart_String(t *testing.T) {
	for _, c := range []struct {
		msg  string
		part Part
		exp  string
	}{
		{"", Part{Tok: Error, Val: "error message"}, `Error[0,0]: "error message"`},
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
