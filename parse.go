package timesheet

import (
	"fmt"
)

func ParseString(body, name string) {
	l := &lexer{
		report: make(chan part),
	}
	go l.lex()
	for {
		select {
		case p, more := <-l.report:
			fmt.Println(p)
			if !more {
				l.report = nil
			}
		}
		if l.report == nil {
			break
		}
	}
}
