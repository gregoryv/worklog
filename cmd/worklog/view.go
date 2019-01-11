package main

import timesheet "github.com/gregoryv/go-timesheet"

type ReportView struct {
	Sheets         []SheetView
	Expected       string
	Reported       string
	ReportedIndent string
	Diff           string
	Tags           []TagView
}

type SheetView struct {
	Period   string
	Expected string
	Reported string
	Diff     string
	Tags     []timesheet.Tagged
}

type TagView struct {
	Duration string
	Tag      string
}

func convertToTagView(tags []timesheet.Tagged) []TagView {
	view := make([]TagView, len(tags))
	for i, t := range tags {
		view[i] = TagView{
			Duration: hhmm(t.Duration),
			Tag:      t.Tag,
		}
	}
	return view
}
