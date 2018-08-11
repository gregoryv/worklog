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
		{"", Part{tok: Unknown, val: "jib"}, `Unknown: "jib"`},
		{"", Part{tok: Number, val: "1"}, `Number: "1"`},
		{"Undefined, run 'go generate'", Part{tok: Token(-1), val: ""}, `Token(-1): ""`},
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
