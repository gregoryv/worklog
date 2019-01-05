package timesheet

import (
	"testing"
)

func TestPart_Defined(t *testing.T) {
	cases := []struct {
		part Part
		exp  bool
	}{
		{Part{}, false},
		{Part{Tok: Error}, true},
	}
	for _, c := range cases {
		got := c.part.Defined()
		exp := c.exp
		if got != exp {
			t.Errorf("%v, expected %v", got, exp)
		}
	}
}

func TestPart_Equals(t *testing.T) {
	cases := []struct {
		a, b Part
		exp  bool
	}{
		{
			Part{Tok: Year, Val: "1", Pos: Position{1, 1}},
			Part{Tok: Year, Val: "1", Pos: Position{1, 1}},
			true,
		},
	}
	for _, c := range cases {
		got := c.a.Equals(c.b)
		exp := c.exp
		if got != exp {
			t.Errorf("%v, expected %v", got, exp)
		}
	}
}

func TestPart_Errorf(t *testing.T) {
	p := Part{Tok: Year, Val: "12x3"}
	got := p.Errorf("invalid %s", "12x").Error()
	exp := "invalid 12x"
	if got != "invalid 12x" {
		t.Errorf("%q, expected %q", got, exp)
	}
	if p.Tok != Error {
		t.Errorf("%v, expected %v", p.Tok, Error)
	}
	if p.Val != exp {
		t.Errorf("%q, expected %q", p.Val, exp)
	}
}

func TestPart_String(t *testing.T) {
	for _, c := range []struct {
		msg  string
		part Part
		exp  string
	}{
		{"", Part{Tok: Error, Val: "error message"}, `Error[0,0]: "error message"`},
		{"", Part{Tok: Year, Val: "1"}, `Year[0,0]: "1"`},
		{"Undefined, run 'go generate'", Part{Tok: Token(-1), Val: ""},
			`Token(-1)[0,0]: ""`},
	} {
		got := c.part.String()
		exp := c.exp
		if got != exp {
			t.Errorf("%q, expected %q", got, exp)
		}
	}
}

func TestNewPart(t *testing.T) {
	got := NewPart()
	if &got == nil {
		t.Fail()
	}
}
