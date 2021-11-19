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
	"text/template"
	"time"

	"github.com/gregoryv/cmdline"
	timesheet "github.com/gregoryv/go-timesheet"
)

func main() {
	var (
		cli     = cmdline.NewBasicParser()
		verbose = cli.Flag("--verbose")
		origin  = cli.Option("--origin",
			"Original timesheets, for comparing reported").String("")
	)
	cli.Parse()

	cmd := Worklog{
		out:     os.Stdout,
		verbose: verbose,
		origin:  origin,
	}
	filePaths := cli.Args()
	if len(filePaths) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	err := cmd.run(filePaths)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type Worklog struct {
	out io.Writer

	verbose bool
	origin  string
}

func (me *Worklog) run(filePaths []string) error {
	expect := timesheet.NewReport()
	report := timesheet.NewReport()
	for _, tspath := range filePaths {
		if me.verbose {
			fmt.Fprintln(os.Stderr, tspath)
		}
		sheet, err := timesheet.Load(tspath)
		if err != nil {
			return err
		}
		report.Append(sheet)
		if me.origin != "" {
			tspath := path.Join(me.origin, path.Base(tspath))
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

	return renderText(me.out, view)
}

func hhmm(dur time.Duration) string {
	return fmt.Sprintf("%7s", timesheet.FormatHHMM(dur))
}

// diff returns difference between reported and expected duration
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

func ConvertToTagView(tags []timesheet.Tagged) []TagView {
	view := make([]TagView, len(tags))
	for i, t := range tags {
		view[i] = TagView{
			Duration: hhmm(t.Duration),
			Tag:      t.Tag,
		}
	}
	return view
}

func renderText(w io.Writer, view *ReportView) error {
	t, err := template.New("default").Parse(`{{range .Sheets}}{{.Period}} {{.Reported}} {{.Diff}} {{range .Tags}} ({{.}}){{end}}
{{end}}
{{printf "%22s" .ReportedIndent}} {{.Diff}}
{{range .Tags}}{{printf "%30s" ""}} {{.Duration}} {{.Tag}}
{{end}}`)

	if err != nil {
		return err
	}
	return t.Execute(w, view)
}
