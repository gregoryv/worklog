package timesheet

import (
	. "github.com/gregoryv/qual"
	"testing"
)

func TestPart_Errorf(t *testing.T) {
	p := Part{Tok: Number, Val: "12x3"}
	got := p.Errorf("invalid %s", "12x").Error()
	Assert(t, Vars{got, p.Val, p.Tok},
		got == "invalid 12x",
		p.Val == "invalid 12x",
		p.Tok == Error,
	)
}

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
		Assert(t, Vars{c.msg, c.exp, got},
			c.exp == got,
		)
	}
}

func TestNewPart(t *testing.T) {
	part := NewPart()
	if part == nil {
		t.Fail()
	}
}
