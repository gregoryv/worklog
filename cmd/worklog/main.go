package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"time"

	timesheet "github.com/gregoryv/go-timesheet"
)

func main() {
	employee := flag.String("employee", "", "Name of Employee")
	html := flag.String("html", "", "Html template")
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

	expect := timesheet.NewReport()
	report := timesheet.NewReport()
	report.Employee = *employee
	for _, tspath := range filePaths {
		sheet, err := timesheet.Load(tspath)
		fatal(err, tspath)
		report.Append(sheet)
		if origin != "" {
			tspath := path.Join(origin, path.Base(tspath))
			esheet, err := timesheet.Load(tspath)
			expect.Append(esheet)
			if err != nil {
				// log perhaps
			}
		}
	}
	view := &ReportView{
		Expected: hhmm(expect.Reported()),
		Reported: hhmm(report.Reported()),
		Diff:     diff(report.Reported(), expect.Reported()),
		Tags:     report.Tags(),
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

	if *html != "" {
		err := renderHtml(os.Stdout, view, *html)
		fatal(err, *html)
		return
	}

	err := renderText(os.Stdout, view, *textTemplate)
	fatal(err, *textTemplate)
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

func fatal(err error, path string) {
	if err != nil {
		fmt.Println(path, err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Printf("Usage: %s TIMESHEET...\n", os.Args[0])
	flag.PrintDefaults()
}
