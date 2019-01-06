package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	timesheet "github.com/gregoryv/go-timesheet"
)

func main() {
	employee := flag.String("e", "", "Name of Employee")
	html := flag.String("html", "", "Html template")
	textTemplate := flag.String("text", "", "Text template")
	flag.Usage = usage
	flag.Parse()

	filePaths := flag.Args()
	if len(filePaths) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	p := timesheet.NewParser()
	view := NewView()
	view.Employee = *employee
	for _, path := range filePaths {
		body, err := ioutil.ReadFile(path)
		fatal(err, path)

		sheet, err := p.Parse(body)
		fatal(err, path)
		view.Sheets = append(view.Sheets, *sheet)
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
