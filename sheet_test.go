package timesheet

import (
	"bytes"
	"testing"
	"time"
)

func TestRender(t *testing.T) {
	w := bytes.NewBufferString("")
	Render(w, 2019, 1, 8)
	p := NewParser()
	sheet, err := p.Parse(w.Bytes())
	if err != nil {
		t.Errorf("%v\n%v", err, w.String())
	}
	if sheet.Period != "2019 January" {
		t.Errorf("Wrong period: %s", sheet.Period)
	}
}

func TestParser_Parse(t *testing.T) {
	p := NewParser()
	sheet, err := p.Parse([]byte(`2018 January
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
	exp := "1977 January 0:00 reported (1:00 flex)"
	if got != exp {
		t.Errorf("\n%q, expected\n%q", got, exp)
	}
}
