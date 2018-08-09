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

func atYear(l *lexer) state {
	l.nextAs(Year)
	return atMonth
}

func atMonth(l *lexer) state {
	l.nextAs(Month)
	l.stream.Next() // Skip \n
	return atSeparator(l)
}

func atSeparator(l *lexer) state {
	p := part{
		tag: Separator,
		val: l.scanAll("-"),
	}
	l.report <- p
	return nil
}

func (l *lexer) scanAll(valid string) string {
	var s string
	for {
		r := l.stream.Next()
		if !strings.ContainsRune(valid, r) {
			break
		}
		s += string(r)
	}
	return s
}

func (l *lexer) nextAs(tag tag) {
	p := part{tag: tag}
	l.stream.Scan()
	p.val = l.stream.TokenText()
	l.report <- p
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
