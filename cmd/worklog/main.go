package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
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

	// todo how to use the original for expected reported time
	if origin != "" {
		originalPaths, err := filepath.Glob(path.Join(origin, "*.timesheet"))
		fatal(err, origin)
		for _, path := range originalPaths {
			fmt.Println(filepath.Base(path))
		}
	}

	view := NewView()
	view.Employee = *employee
	for _, path := range filePaths {
		sheet, err := timesheet.Load(path)
		fatal(err, path)
		view.Append(sheet)
	}
	if *html != "" {
		err := renderHtml(os.Stdout, view, *html)
		fatal(err, *html)
		return
	}

	err := renderText(os.Stdout, view, *textTemplate)
	fatal(err, *textTemplate)
}

type View struct {
	Employee string
	Sheets   []timesheet.Sheet
}

func NewView() *View {
	return &View{
		Sheets: make([]timesheet.Sheet, 0),
	}
}

func (view *View) Append(sheet *timesheet.Sheet) {
	view.Sheets = append(view.Sheets, *sheet)
}

func (view *View) Reported() string {
	var reported time.Duration
	for _, sheet := range view.Sheets {
		reported += sheet.Reported.Duration
	}
	return timesheet.FormatHHMM(reported)
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
