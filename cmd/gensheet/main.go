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
	year := time.Now().Year()
	flag.IntVar(&year, "y", year, "Year, four digits")
	month := int(time.Now().Month())
	flag.IntVar(&month, "m", month, "Month, 1-12")
	out := ""
	flag.StringVar(&out, "o", out, "Save timesheets in directory, use with -m -1")
	flag.Parse()

	if month == -1 {
		m := 1
		if out == "" {
			fatal(fmt.Errorf("Missing -o, see help"))
		}
		err := os.MkdirAll(out, 0744)
		fatal(err)
		for {
			filepath := path.Join(out, fmt.Sprintf("%v%02v.timesheet", year, m))
			file, err := os.Create(filepath)
			fatal(err)
			timesheet.Render(file, year, time.Month(m), 8)
			file.Close()
			m++
			if m == 13 {
				break
			}
		}
		return
	}
	timesheet.Render(os.Stdout, year, time.Month(month), 8)
}

func fatal(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
