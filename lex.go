package timesheet

import (
	"fmt"
	"strings"
	"text/scanner"
)

// lexer uses a scanner to find parts within a stream
type lexer struct {
	report chan part
	stream scanner.Scanner
}

// state funcs are used to scan for parts from specific context
type state func(*lexer) state

func start(l *lexer) state {
	r := l.stream.Peek()
	// If comment

	// or Year Month row is
	// If starts with digit atRow
	if strings.ContainsRune("0123456789", r) {
		return atYear
	}
	return nil
}

func atYear(l *lexer) state {
	p := part{tag: Year}
	l.stream.Scan()
	p.val = l.stream.TokenText()
	l.report <- p
	return nil
}

// lex runs until report is closed
func (l *lexer) lex() {
	for state := start; state != nil; {
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
