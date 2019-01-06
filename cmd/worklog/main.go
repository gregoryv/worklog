package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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
		fmt.Print(path, " ")
		body, err := ioutil.ReadFile(path)
		fatal(err)

		reported, err := p.SumReported(body)
		fatal(err)
		tag := timesheet.Tagged{reported, "reported"}
		fmt.Print(tag, " ")

		tagged, err := p.SumTagged(body)
		fatal(err)
		for _, tag := range tagged {
			fmt.Print("(", tag, ") ")
		}
		fmt.Println()
	}
}

func srender(tag timesheet.Tagged) string {
	parts := strings.Split(tag.String(), " ")
	return fmt.Sprintf("%7s %s", parts[0], parts[1])
}

func fatal(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
