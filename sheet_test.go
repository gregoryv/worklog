package timesheet

import "testing"

func TestParser_Parse(t *testing.T) {
	p := NewParser()
	sheet, err := p.Parse([]byte(""))
	if sheet == nil {
		t.Errorf("Expected a sheet")
	}
	if err == nil {
		t.Error("Expected error")
	}
}
