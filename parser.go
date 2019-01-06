package timesheet

import (
	"fmt"
	"strconv"
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

func (par *Parser) Sum(body []byte) (hh, mm int) {
	lex := NewLexer(string(body))
	out := lex.Run()
	var sum int
	var inTag bool
	for {
		p, more := <-out
		if !more {
			break
		}
		switch p.Tok {
		case LeftParenthesis:
			inTag = true
		case RightParenthesis:
			inTag = false
		case Hours:
			if !inTag {
				h, _ := strconv.Atoi(p.Val)
				sum += (60 * h)
			}
		case Minutes:
			if !inTag {
				m, _ := strconv.Atoi(p.Val)
				sum += m
			}
		case Error:
			fmt.Println("Error", p.Val)
		}
	}
	hh = sum / 60
	mm = sum - 60*hh
	return
}
