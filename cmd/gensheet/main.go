package main

import (
	"flag"
	"os"
	"time"

	timesheet "github.com/gregoryv/go-timesheet"
)

func main() {
	year := time.Now().Year()
	flag.IntVar(&year, "y", year, "Year, four digits")
	month := int(time.Now().Month())
	flag.IntVar(&month, "m", month, "Month, 1-12")
	flag.Parse()

	timesheet.Render(os.Stdout, year, time.Month(month), 8)
}
