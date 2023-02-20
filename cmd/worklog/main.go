// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"text/template"
	"time"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/worklog"
)

func main() {
	var (
		cli = cmdline.NewBasicParser()
		cmd = Worklog{
			out:     os.Stdout,
			verbose: cli.Flag("--verbose"),

			origin: cli.Option("--originals", "For calculating flex").String(""),

			filenames: cli.NamedArg("FILES...").Strings(),
		}
	)
	u := cli.Usage()
	u.Example(
		"Generate report",
		"    $ worklog *.timesheet",
	)
	u.Example(
		"Report compared to originals",
		"    $ worklog --originals ./path/to/dir/ *.timesheet",
		"",
	)
	cli.Parse()

	if err := cmd.Run(); err != nil {
		log.SetFlags(0)
		log.Fatal(err)
	}
}

type Worklog struct {
	out io.Writer

	verbose bool
	origin  string

	filenames []string
}

// Run generates a comparison of timesheets to the original as a
// condensed report
func (me *Worklog) Run() error {
	expect := worklog.NewReport()
	report := worklog.NewReport()
	for _, filename := range me.filenames {
		if me.verbose {
			log.Println(filename)
		}
		sheet, err := worklog.Load(filename)
		if err != nil {
			return err
		}
		report.Append(sheet)
		// calculate expected timesheet
		if me.origin != "" {
			// original filename must match the given filename
			filename := path.Join(me.origin, path.Base(filename))
			esheet, err := worklog.Load(filename)
			if err == nil {
				expect.Append(esheet)
			}
		}
	}

	view := &ReportView{
		Expected:       hhmm(expect.Reported()),
		Reported:       hhmm(report.Reported()),
		ReportedIndent: fmt.Sprintf("%22s", hhmm(report.Reported())),
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
		// todo look for flex tag
		exp, _ := expect.FindByPeriod(sheet.Period)
		if exp != nil {
			view.Expected = worklog.FormatHHMM(exp.Reported.Duration)
			view.Diff = diff(sheet.Reported.Duration, exp.Reported.Duration)
		}
		sheetViews = append(sheetViews, view)
	}
	view.Sheets = sheetViews

	return renderText(me.out, view)
}

func hhmm(dur time.Duration) string {
	return fmt.Sprintf("%7s", worklog.FormatHHMM(dur))
}

// diff returns difference between reported and expected duration
func diff(rep, exp time.Duration) string {
	diff := rep - exp
	var d string
	switch {
	case diff < 0:
		d = worklog.FormatHHMM(diff)
	case diff > 0:
		d = "+" + worklog.FormatHHMM(diff)
	}
	return fmt.Sprintf("%7s", d)
}

type ReportView struct {
	Sheets         []SheetView
	Expected       string
	Reported       string
	ReportedIndent string
	Diff           string
	Tags           []TagView
}

// SheetView is the view model for each montly line
type SheetView struct {
	Period   string
	Expected string
	Reported string
	Diff     string
	Tags     []worklog.Tagged
}

type TagView struct {
	Duration string
	Tag      string
}

func ConvertToTagView(tags []worklog.Tagged) []TagView {
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
{{ .ReportedIndent}} {{.Diff}}
{{range .Tags}}{{printf "%30s" ""}} {{.Duration}} {{.Tag}}
{{end}}`)

	if err != nil {
		return err
	}
	return t.Execute(w, view)
}
