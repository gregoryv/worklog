package timesheet

import (
	"fmt"
	"time"
)

type Tagged struct {
	Duration time.Duration
	Tag      string
}

func (tagged Tagged) String() string {
	dur := tagged.Duration
	hh := dur.Truncate(time.Hour)
	var operator time.Duration = 1
	if hh < 0 {
		operator = -1
	}
	mm := (dur - hh) * operator
	return fmt.Sprintf("%v:%02v %s", hh.Hours(), mm.Minutes(), tagged.Tag)
}

func (par *Parser) SumTagged(body []byte) ([]Tagged, error) {
	sheet, err := par.Parse(body)
	return sheet.Tags, err
}

type byTag []Tagged

func (by byTag) Len() int           { return len(by) }
func (by byTag) Less(i, j int) bool { return by[i].Tag < by[j].Tag }
func (by byTag) Swap(i, j int) {
	by[i], by[j] = by[j], by[i]
}
