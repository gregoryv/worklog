package timesheet

import (
	"fmt"
	"strings"
	"text/scanner"
)

func Example() {
	report := make(chan part)
	var s scanner.Scanner
	s.Init(strings.NewReader("2018 August"))
	l := &lexer{
		report: report,
		stream: s,
	}
	// print all parts as they come in
	done := make(chan bool)
	go func() {
		for {
			part, more := <-report
			if !more {
				done <- true
				return
			}
			fmt.Println(part)
		}
	}()
	l.lex()
	<-done
	// output:
	// Year: 2018
	// Month: August
}
