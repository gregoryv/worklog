package main

import timesheet "github.com/gregoryv/go-timesheet"

type ReportView struct {
	Sheets   []SheetView
	Expected string
	Reported string
	Diff     string
	Tags     []timesheet.Tagged
}

type SheetView struct {
	Period   string
	Expected string
	Reported string
	Tags     []timesheet.Tagged
}
