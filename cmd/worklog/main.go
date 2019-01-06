package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	timesheet "github.com/gregoryv/go-timesheet"
)

func main() {
	flag.Parse()

	filePaths := flag.Args()
	if len(filePaths) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	p := timesheet.NewParser()
	for _, path := range filePaths {
		body, err := ioutil.ReadFile(path)
		fatal(err, path)

		sheet, err := p.Parse(body)
		fatal(err, path)
		fmt.Println(sheet)
	}
}

func fatal(err error, path string) {
	if err != nil {
		fmt.Println(path, err)
		os.Exit(1)
	}
}
