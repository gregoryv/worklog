package timesheet

import (
	"fmt"
	"testing"
)

func ExamplePos_String() {
	p := NewPos()
	fmt.Println(p)
	//output:
	//0,0
}

func TestNewPos(t *testing.T) {
	if p := NewPos(); p == nil {
		t.Fail()
	}
}
