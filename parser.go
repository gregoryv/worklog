package timesheet

import (
	"fmt"
	"strconv"
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
	lex := NewLexer(string(body))
	out := lex.Run()
	inTag := false
	for {
		p, more := <-out
		switch p.Tok {
		case LeftParenthesis:
			inTag = true
		case RightParenthesis:
			inTag = false
		case Hours:
			if !inTag {
				h, _ := strconv.Atoi(p.Val)
				dur += time.Duration(h*60*60) * time.Second
			}
		case Minutes:
			if !inTag {
				m, _ := strconv.Atoi(p.Val)
				dur += time.Duration(m*60) * time.Second
			}
		case Error:
			return 0, fmt.Errorf("%s", p)
		}
		if !more {
			break
		}
	}
	return
}
