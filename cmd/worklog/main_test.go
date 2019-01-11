package main

import (
	"os/exec"
	"testing"

	"github.com/gregoryv/asserter"
)

func TestFeature(t *testing.T) {
	out, err := exec.Command("worklog", "-origin", "../../assets/orig2018",
		"../../assets/201801.timesheet").CombinedOutput()
	if err != nil {
		t.Fatal(err, string(out))
	}
	got := string(append([]byte("\n"), out...))
	exp := `
2018 January    179:30   +7:30  (8:00 semester)

 179:30   +7:30
8:00 semester
`
	assert := asserter.New(t)
	assert().Equals(got, exp)
}
