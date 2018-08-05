package timesheet

import (
	"io"
	"text/scanner"
)

type Parser struct {
	Debug Logger
}

func NewParser() *Parser {
	var noop NilLogger
	return &Parser{
		Debug: noop,
	}
}

func (p *Parser) Parse(in io.Reader, filename string) (err error) {
	return p.parse(in, filename, p.Debug)
}

func (p *Parser) parse(in io.Reader, filename string, debug Logger) (err error) {
	var s scanner.Scanner
	s.Init(in)
	s.Filename = filename
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		debug.Logf("%s: %s\n", s.Position, s.TokenText())
	}
	return
}
