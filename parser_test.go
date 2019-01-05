package timesheet

import (
	"testing"
)

func ExampleParser_Dump() {
	par := NewParser()
	sheet := `2018 January
----------
1  1 Mon 8 (4 semester) was in thailand (2 pool)`
	par.Dump([]byte(sheet))
	// output:
	// Year[1,1]: "2018"
	// Month[1,6]: "January"
	// Separator[2,1]: "----------"
	// Week[3,1]: "1"
	// Date[3,4]: "1"
	// Day[3,6]: "Mon"
	// Hours[3,10]: "8"
	// Note[3,11]: " "
	// LeftParenthesis[3,12]: "("
	// Hours[3,13]: "4"
	// Tag[3,14]: "semester"
	// RightParenthesis[3,23]: ")"
	// Note[3,25]: "was in thailand "
	// LeftParenthesis[3,41]: "("
	// Hours[3,42]: "2"
	// Tag[3,43]: "pool"
	// RightParenthesis[3,48]: ")"
}

func TestParser_Sum(t *testing.T) {
	par := NewParser()
	sheet := `2018 January
----------
1  1 Mon 8 (4 semester) thailand (2 pool)`
	gotHH, gotMM := par.Sum([]byte(sheet))
	expHH, expMM := 8, 0
	if gotHH != expHH {
		t.Errorf("%q, expected %q", gotHH, expHH)
	}
	if gotMM != expMM {
		t.Errorf("%q, expected %q", gotMM, expMM)
	}

}
