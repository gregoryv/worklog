package timesheet

import (
	"fmt"
	"testing"
)

func TestParser_SumTagged(t *testing.T) {
	sheet := []byte(`2018 January
----------
1  1 Mon 8 (4 vacation) was in thailand (2:30 pool)
   2 Tue 4:10 (4 vacation) was in thailand (-1 pool)`)

	got, _ := NewParser().SumTagged(sheet)
	if len(got) != 2 {
		t.Errorf("%v, expected %v", got, 2)
	}
	{
		got := fmt.Sprintf("%v", got)
		exp := "[01:30 pool 08:00 vacation]"
		if got != exp {
			t.Errorf("%v, expected %v", got, exp)
		}
	}
}

func TestParser_SumTagged_errors(t *testing.T) {
	sheet := []byte(`2018 January
----------
1  1 Mon 8 (4 vacation) (2 pool`)

	got, err := NewParser().SumTagged(sheet)
	exp := 0
	if len(got) != exp {
		t.Errorf("%v, expected %v", len(got), exp)
	}
	if err == nil {
		t.Errorf("Expected an error from SumTagged\n%v\n", string(sheet))
	}
}
