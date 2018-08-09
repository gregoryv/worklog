package timesheet

import (
	"fmt"
	"text/scanner"
)

// lexer uses a scanner to find parts within a stream
type lexer struct {
	report chan part
	stream scanner.Scanner
}

// state funcs are used to scan for parts from specific context
type state func(*lexer) state

func atYear(l *lexer) state {
	p := part{tag: Year}
	l.stream.Scan()
	p.val = l.stream.TokenText()
	l.report <- p
	return atMonth
}

func atMonth(l *lexer) state {
	p := part{tag: Month}
	l.stream.Scan()
	p.val = l.stream.TokenText()
	l.report <- p
	return nil
}

// lex runs until report is closed
func (l *lexer) lex() {
	for state := atYear; state != nil; {
		state = state(l)
	}
	close(l.report)
}

type part struct {
	tag tag
	val string
}

func (p part) String() string {
	return fmt.Sprintf("%s: %v", p.tag, p.val)
}
