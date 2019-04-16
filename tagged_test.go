// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package timesheet

import (
	"fmt"
	"testing"
)

func TestSumTagged(t *testing.T) {
	sheet := []byte(`2018 January
----------
1  1 Mon 8 (4 vacation) was in thailand (2:30 pool)
   2 Tue 4:10 (4 vacation) was in thailand (-1 pool) (-2:10 flex)`)

	tagged, _ := SumTagged(sheet)
	got := fmt.Sprintf("%v", tagged)
	exp := "[-2:10 flex 1:30 pool 8:00 vacation]"
	if got != exp {
		t.Errorf("%v, expected %v", got, exp)
	}
}

func TestSumTagged_errors(t *testing.T) {
	sheet := []byte(`2018 January
----------
1  1 Mon 8 (4 vacation) (2 pool`)

	got, err := SumTagged(sheet)
	exp := 0
	if len(got) != exp {
		t.Errorf("%v, expected %v", len(got), exp)
	}
	if err == nil {
		t.Errorf("Expected an error from SumTagged\n%v\n", string(sheet))
	}
}
