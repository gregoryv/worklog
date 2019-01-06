package timesheet

import (
	"testing"
)

func ExampleParser_Dump() {
	NewParser().Dump([]byte(`2018 January
----------
1  1 Mon 8:10 (4 semester) was in thailand (2 pool)
`))
	// output:
	// Year[1,1]: "2018"
	// Month[1,6]: "January"
	// Separator[2,1]: "----------"
	// Week[3,1]: "1"
	// Date[3,4]: "1"
	// Day[3,6]: "Mon"
	// Hours[3,10]: "8"
	// Colon[3,11]: ":"
	// Minutes[3,12]: "10"
	// LeftParenthesis[3,15]: "("
	// Hours[3,16]: "4"
	// Tag[3,18]: "semester"
	// RightParenthesis[3,26]: ")"
	// Note[3,28]: "was in thailand "
	// LeftParenthesis[3,44]: "("
	// Hours[3,45]: "2"
	// Tag[3,47]: "pool"
	// RightParenthesis[3,51]: ")"
}

var parserTestSheet = `2018 January
----------
1  1 Mon 8 (4 vacation) was in thailand (2:30 pool)
   2 Tue 4:10 (4 vacation) was in thailand`

func TestParser_Sum_hours(t *testing.T) {
	got, _ := NewParser().Sum([]byte(parserTestSheet))
	exp := 12
	if got != exp {
		t.Errorf("%v, expected %v", got, exp)
	}
}

func TestParser_Sum_min(t *testing.T) {
	_, got := NewParser().Sum([]byte(parserTestSheet))
	exp := 10
	if got != exp {
		t.Errorf("%v, expected %v", got, exp)
	}
}
