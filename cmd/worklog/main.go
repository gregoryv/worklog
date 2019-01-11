package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"

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

	// todo move view into a timesheet.Report | Book | Ledger
	origReport := timesheet.NewReport()
	// todo how to use the original for expected reported time
	if origin != "" {
		originalPaths, err := filepath.Glob(path.Join(origin, "*.timesheet"))
		fatal(err, origin)
		for _, path := range originalPaths {
			sheet, err := timesheet.Load(path)
			fatal(err, path)
			origReport.Append(sheet)
		}
	}

	report := timesheet.NewReport()
	report.Employee = *employee
	for _, path := range filePaths {
		sheet, err := timesheet.Load(path)
		fatal(err, path)
		report.Append(sheet)
	}
	if *html != "" {
		err := renderHtml(os.Stdout, report, *html)
		fatal(err, *html)
		return
	}

	err := renderText(os.Stdout, report, *textTemplate)
	fatal(err, *textTemplate)
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
