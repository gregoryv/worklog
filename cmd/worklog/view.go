package main

import (
	"fmt"

	timesheet "github.com/gregoryv/go-timesheet"
)

type View struct {
	timesheet.Report
}

func (r *View) SumReported() string {
	var tags string
	for _, tag := range r.Tags() {
		tags += fmt.Sprintf("%22s %s\n", timesheet.FormatHHMM(tag.Duration), tag.Tag)
	}

	return fmt.Sprintf("%-14s %7s\n%s", "Sum:",
		timesheet.FormatHHMM(r.Reported()), tags)
}
