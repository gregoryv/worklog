// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package timesheet

import (
	"bytes"
	"testing"
	"time"
)

func TestLoad(t *testing.T) {
	_, err := Load("testdata/201506.timesheet")
	if err != nil {
		t.Errorf("Load failed: %v", err)
	}
	_, err = Load("nosuchfile")
	if err == nil {
		t.Error("Expected Load to fail")
	}
}

func TestRender(t *testing.T) {
	w := bytes.NewBufferString("")
	Render(w, 2019, 1, 8)
	sheet, err := Parse(w.Bytes())
	if err != nil {
		t.Errorf("%v\n%v", err, w.String())
	}
	if sheet.Period != "2019 January" {
		t.Errorf("Wrong period: %s", sheet.Period)
	}
}

func TestParse(t *testing.T) {
	sheet, err := Parse([]byte(`2018 January
----------
1  1 Sun
   2 Mon +8 (4 vacation) was in thailand (+2:30 pool)
   3 Tue 4:10 (4 vacation) was in thailand
   4 Wed -1`))
	if sheet == nil {
		t.Errorf("Expected a sheet")
	}
	if err != nil {
		t.Error(err)
	}
	exp := "2018 January"
	if sheet.Period != exp {
		t.Errorf("\n%q, expected\n%q", sheet.Period, exp)
	}
}

func TestSheet_String(t *testing.T) {
	sheet := NewSheet()
	sheet.Period = "1977 January"
	sheet.Tags = []Tagged{{time.Hour, "flex"}}
	got := sheet.String()
	exp := "1977 January      0:00 (1:00 flex)"
	if got != exp {
		t.Errorf("\n%q, expected\n%q", got, exp)
	}
}
