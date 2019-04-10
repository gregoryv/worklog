// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.
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
	month := int(time.Now().Month())
	hours := 8
	flag.IntVar(&year, "y", year, "Year, four digits")
	flag.IntVar(&month, "m", month, "Month, 1-12")
	flag.IntVar(&hours, "w", hours, "Default workhours")
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
			timesheet.Render(file, year, time.Month(m), hours)
			file.Close()
			m++
			if m == 13 {
				break
			}
		}
		return
	}
	timesheet.Render(os.Stdout, year, time.Month(month), hours)
}

func fatal(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
