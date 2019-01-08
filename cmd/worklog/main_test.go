package main

import (
	"os/exec"
	"testing"

	"github.com/gregoryv/asserter"
)

func TestFeature(t *testing.T) {
	out, err := exec.Command("worklog", "../../201506.timesheet").CombinedOutput()
	if err != nil {
		t.Fatal(err, string(out))
	}
	got := string(append([]byte("\n"), out...))
	exp := `
2015 June       174:30 (7:00 conference) (-1:30 flex) (1:00 travel)

Sum:            174:30 (7:00 conference) (-1:30 flex) (1:00 travel)
`
	assert := asserter.New(t)
	assert().Equals(got, exp)
}
