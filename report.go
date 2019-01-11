package timesheet

import (
	"fmt"
	"sort"
	"time"
)

type Report struct {
	Employee string
	Sheets   []Sheet

	reported time.Duration
	tagged   map[string]time.Duration
}

func NewReport() *Report {
	return &Report{
		Sheets:   make([]Sheet, 0),
		reported: 0,
	}
}

func (report *Report) Append(sheet *Sheet) (sheetCount int, err error) {
	report.Sheets = append(report.Sheets, *sheet)
	report.reported += sheet.Reported.Duration
	return len(report.Sheets), nil
}

func (report *Report) Reported() time.Duration {
	return report.reported
}

func (report *Report) SumReported() string {
	var tags string
	for _, tag := range report.SumTagged() {
		tags += fmt.Sprintf("%22s %s\n", FormatHHMM(tag.Duration), tag.Tag)
	}

	return fmt.Sprintf("%-14s %7s\n%s", "Sum:",
		FormatHHMM(report.reported), tags)
}

func (report *Report) SumTagged() []Tagged {
	sum := make(map[string]time.Duration)
	for _, sheet := range report.Sheets {
		for _, tag := range sheet.Tags {
			if _, exists := sum[tag.Tag]; !exists {
				sum[tag.Tag] = 0
			}
			sum[tag.Tag] += tag.Duration
		}
	}
	tags := make([]Tagged, 0)
	for k, v := range sum {
		tags = append(tags, Tagged{v, k})
	}
	sort.Sort(ByTag(tags))
	return tags
}
