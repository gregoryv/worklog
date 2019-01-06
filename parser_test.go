package timesheet

import (
	"testing"
	"time"
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

func ExampleParser_Dump_bad() {
	NewParser().Dump([]byte(`2018 nosuchmonth`))
	// output:
	// Year[1,1]: "2018"
	// Error[1,6]: "invalid Month"
	// Error[1,6]: "invalid Separator"
	// Error[1,6]: "invalid Week"
	// Error[1,6]: "invalid Date"
	// Error[1,6]: "invalid Day"
	// Error[1,6]: "invalid Hours"
	// Note[1,6]: "nosuchmonth"
}

func TestParser_SumReported(t *testing.T) {
	sheet := []byte(`2018 January
----------
1  1 Sun
   2 Mon +8 (4 vacation) was in thailand (+2:30 pool)
   3 Tue 4:10 (4 vacation) was in thailand
   4 Wed -1`)

	got, err := NewParser().SumReported(sheet)
	if err != nil {
		t.Fatalf("%s\n%s", err, string(sheet))
	}
	exp := time.Duration((11*60 + 10) * time.Minute)
	if got != exp {
		t.Errorf("%v, expected %v", got, exp)
	}
}

func TestParser_SumReported_error(t *testing.T) {
	sheet := `2018 January
----------
1  1 Mon 8 (4 vacation) was in thailand (2:30 pool)
   2 Tu 4:10 (4 vacation)` // should be Tue not Tu

	got, err := NewParser().SumReported([]byte(sheet))
	exp := time.Duration(8 * time.Hour) // only first day has been counted
	if got != exp {
		t.Errorf("\n%v\n%v, expected %v", sheet, got, exp)
	}
	if err == nil {
		t.Errorf("Expected error from SumReported\n%v\n", string(sheet))
	}
}
