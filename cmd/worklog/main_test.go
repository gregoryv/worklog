package main

import (
	"os/exec"
	"testing"
)

func TestFeature(t *testing.T) {
	out, err := exec.Command("worklog", "../../201506.timesheet").Output()
	if err != nil {
		t.Fatal(err)
	}
	got := string(out)
	exp := "2015 June 174:30 reported (7:00 conference) (-1:30 flex) (1:00 travel)\n"
	if got != exp {
		t.Errorf("\n%q, expected\n%q", got, exp)
	}
}
