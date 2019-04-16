// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	timesheet "github.com/gregoryv/go-timesheet"
)

func main() {
	textTemplate := flag.String("text", "", "Text template")
	origin := ""
	flag.StringVar(&origin, "origin", origin, "Original timesheets, eg. for comparing reported")
	flag.Usage = usage
	flag.Parse()

	filePaths := flag.Args()
	if len(filePaths) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	err := writeText(os.Stdout, *textTemplate, origin, filePaths)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func writeText(w io.Writer, textTemplate, origin string, filePaths []string) error {
	expect := timesheet.NewReport()
	report := timesheet.NewReport()
	for _, tspath := range filePaths {
		sheet, err := timesheet.Load(tspath)
		if err != nil {
			return err
		}
		report.Append(sheet)
		if origin != "" {
			tspath := path.Join(origin, path.Base(tspath))
			esheet, err := timesheet.Load(tspath)
			if err == nil {
				expect.Append(esheet)
			}
		}
	}
	view := &ReportView{
		Expected:       hhmm(expect.Reported()),
		Reported:       hhmm(report.Reported()),
		ReportedIndent: fmt.Sprintf("%22s", ""),
		Diff:           diff(report.Reported(), expect.Reported()),
		Tags:           ConvertToTagView(report.Tags()),
	}
	sheetViews := make([]SheetView, 0)
	for _, sheet := range report.Sheets {
		view := SheetView{
			Period:   fmt.Sprintf("%-14s", sheet.Period),
			Reported: hhmm(sheet.Reported.Duration),
			Tags:     sheet.Tags,
		}
		exp, _ := expect.FindByPeriod(sheet.Period)
		if exp != nil {
			view.Expected = timesheet.FormatHHMM(exp.Reported.Duration)
			view.Diff = diff(sheet.Reported.Duration, exp.Reported.Duration)
		}
		sheetViews = append(sheetViews, view)
	}
	view.Sheets = sheetViews

	return renderText(w, view, textTemplate)
}

func hhmm(dur time.Duration) string {
	return fmt.Sprintf("%7s", timesheet.FormatHHMM(dur))
}

func diff(rep, exp time.Duration) string {
	diff := rep - exp
	var d string
	switch {
	case diff < 0:
		d = timesheet.FormatHHMM(diff)
	case diff > 0:
		d = "+" + timesheet.FormatHHMM(diff)
	}
	return fmt.Sprintf("%7s", d)
}

func usage() {
	fmt.Printf("Usage: %s TIMESHEET...\n", os.Args[0])
	flag.PrintDefaults()
}
