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
	"github.com/gregoryv/worklog"
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
	filenames := cli.Args()
	if len(filenames) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	err := cmd.Run(filenames...)
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

func (me *Worklog) Run(filenames ...string) error {
	expect := worklog.NewReport()
	report := worklog.NewReport()
	for _, tspath := range filenames {
		if me.verbose {
			fmt.Fprintln(os.Stderr, tspath)
		}
		sheet, err := worklog.Load(tspath)
		if err != nil {
			return err
		}
		report.Append(sheet)
		if me.origin != "" {
			tspath := path.Join(me.origin, path.Base(tspath))
			esheet, err := worklog.Load(tspath)
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
{{printf "%22s" .ReportedIndent}} {{.Diff}}
{{range .Tags}}{{printf "%30s" ""}} {{.Duration}} {{.Tag}}
{{end}}`)

	if err != nil {
		return err
	}
	return t.Execute(w, view)
}
