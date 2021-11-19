// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package worklog

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
	mm := abs(dur) - abs(hh)
	operator := ""
	if dur < 0 {
		operator = "-"
	}
	hh = abs(hh)
	mm = abs(mm)
	return fmt.Sprintf("%v%v:%02v", operator, hh.Hours(), mm.Minutes())
}

func abs(dur time.Duration) time.Duration {
	if dur < 0 {
		return -1 * dur
	}
	return dur
}

func SumTagged(body []byte) (tags []Tagged, err error) {
	sheet, err := Parse(body)
	if err != nil {
		return
	}
	return sheet.Tags, nil
}

type ByTag []Tagged

func (by ByTag) Len() int           { return len(by) }
func (by ByTag) Less(i, j int) bool { return by[i].Tag < by[j].Tag }
func (by ByTag) Swap(i, j int) {
	by[i], by[j] = by[j], by[i]
}
