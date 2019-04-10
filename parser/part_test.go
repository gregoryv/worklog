package parser

import (
	"fmt"
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

func ExamplePart_String() {
	fmt.Printf("%s\n%s\n%s",
		Part{Tok: token.Error, Val: "error message"},
		Part{Tok: token.Year, Val: "1"},
		Part{Tok: token.Token(-1), Val: ""},
	)
	// output:
	// Error[0,0]: "error message"
	// Year[0,0]: "1"
	// Token(-1)[0,0]: ""
}
