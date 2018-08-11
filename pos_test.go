package timesheet

import (
	"fmt"
	"testing"
)

func TestPos_Next(t *testing.T) {
	cases := []struct {
		line, col int
	}{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	p := NewPos()
	for i, c := range cases {
		line, col := p.Next()
		switch {
		case c.line != line:
			if i == 0 {
				t.Error("Next() should not advance line")
			}
			t.Errorf("Expected line %v, got %v", c.line, line)
		case c.col != col:
			if i == 0 {
				t.Error("Next() should advance column by one")
			}
			t.Errorf("Expected col %v, got %v", c.col, col)
		}
	}
}

func ExamplePos_String() {
	p := NewPos()
	fmt.Println(p)
	//output:
	//1,1
}

func TestNewPos(t *testing.T) {
	if p := NewPos(); p == nil {
		t.Fail()
	}
}
