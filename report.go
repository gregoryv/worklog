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
	tags     map[string]time.Duration
}

func NewReport() *Report {
	return &Report{
		Sheets:   make([]Sheet, 0),
		reported: 0,
		tags:     make(map[string]time.Duration),
	}
}

func (r *Report) Append(sheet *Sheet) (sheetCount int, err error) {
	r.Sheets = append(r.Sheets, *sheet)
	// sum reported and tagged values
	r.reported += sheet.Reported.Duration
	for _, tag := range sheet.Tags {
		if _, exists := r.tags[tag.Tag]; !exists {
			r.tags[tag.Tag] = 0
		}
		r.tags[tag.Tag] += tag.Duration
	}
	return len(r.Sheets), nil
}

func (r *Report) Reported() time.Duration {
	return r.reported
}

// Tags returns a sorted and summarized list of tags
func (r *Report) Tags() []Tagged {
	tags := make([]Tagged, 0)
	for k, v := range r.tags {
		tags = append(tags, Tagged{v, k})
	}
	sort.Sort(ByTag(tags))
	return tags
}

// TODO Move this to text view somehow
func (r *Report) SumReported() string {
	var tags string
	for _, tag := range r.Tags() {
		tags += fmt.Sprintf("%22s %s\n", FormatHHMM(tag.Duration), tag.Tag)
	}

	return fmt.Sprintf("%-14s %7s\n%s", "Sum:",
		FormatHHMM(r.reported), tags)
}
