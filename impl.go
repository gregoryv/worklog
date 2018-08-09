package timesheet

import (
	"fmt"
)

func (p part) String() string {
	return fmt.Sprintf("%v", p.val)
}

func start(l *lexer) state {
	// Comment

	// or Year Month row is
	// If starts with digit atRow
	return atYear
}

func atYear(l *lexer) state {
	return nil
}

// lex runs until report is closed
func (l *lexer) lex() {
	for state := start; state != nil; {
		state = state(l)
	}
	close(l.report)
}

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
