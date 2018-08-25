package timesheet

import (
	. "github.com/gregoryv/qual"
	"testing"
)

func ExampleParser_Dump() {
	par := NewParser()
	sheet := `2018 January
----------
1  1 Mon 8 (4 semester) thailand (2 pool)`
	par.Dump([]byte(sheet))
	//output:
	// Year[1,1]: "2018"
}

func TestParser_Sum(t *testing.T) {
	par := NewParser()
	sheet := `2018 January
----------
1  1 Mon 8 (4 semester) thailand (2 pool)`
	hh, mm := 8, 0
	gotHH, gotMM := par.Sum([]byte(sheet))
	Assert(t, Vars{hh, gotHH, mm, gotMM},
		hh == gotHH,
		mm == gotMM,
	)
}
