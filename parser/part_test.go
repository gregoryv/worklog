package parser

import (
	"testing"

	"github.com/gregoryv/go-timesheet/token"
)

func TestPart_Defined(t *testing.T) {
	cases := []struct {
		part Part
		exp  bool
	}{
		{Part{}, false},
		{Part{Tok: token.Error}, true},
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
			Part{Tok: token.Year, Val: "1", Pos: Position{1, 1}},
			Part{Tok: token.Year, Val: "1", Pos: Position{1, 1}},
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
	p := Part{Tok: token.Year, Val: "12x3"}
	got := p.Errorf("invalid %s", "12x").Error()
	exp := "invalid 12x"
	if got != "invalid 12x" {
		t.Errorf("%q, expected %q", got, exp)
	}
	if p.Tok != token.Error {
		t.Errorf("%v, expected %v", p.Tok, token.Error)
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
		{"", Part{Tok: token.Error, Val: "error message"}, `Error[0,0]: "error message"`},
		{"", Part{Tok: token.Year, Val: "1"}, `Year[0,0]: "1"`},
		{"Undefined, run 'go generate'", Part{Tok: token.Token(-1), Val: ""},
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
