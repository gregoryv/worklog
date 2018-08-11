package timesheet

import (
	"testing"
)

func TestNewPos(t *testing.T) {
	if p := NewPos(); p == nil {
		t.Fail()
	}
}
