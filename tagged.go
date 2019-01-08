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
	str := FormatHHMM(tagged.Duration)
	if tagged.Tag != "" {
		return str + " " + tagged.Tag
	}
	return str
}

func FormatHHMM(dur time.Duration) string {
	hh := dur.Truncate(time.Hour)
	var operator time.Duration = 1
	if hh < 0 {
		operator = -1
	}
	mm := (dur - hh) * operator
	return fmt.Sprintf("%v:%02v", hh.Hours(), mm.Minutes())

}

func (par *Parser) SumTagged(body []byte) ([]Tagged, error) {
	sheet, err := par.Parse(body)
	return sheet.Tags, err
}

type ByTag []Tagged

func (by ByTag) Len() int           { return len(by) }
func (by ByTag) Less(i, j int) bool { return by[i].Tag < by[j].Tag }
func (by ByTag) Swap(i, j int) {
	by[i], by[j] = by[j], by[i]
}
