package timesheet

import (
	"fmt"
	"time"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (par *Parser) Dump(body []byte) {
	lex := NewLexer(string(body))
	out := lex.Run()
	for {
		p, more := <-out
		if !more {
			break
		}
		fmt.Println(p)
	}
}

func (par *Parser) SumReported(body []byte) (dur time.Duration, err error) {
	sheet, err := par.Parse(body)
	return sheet.Reported.Duration, err
}
