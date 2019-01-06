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

func (par *Parser) SumPreported(body []byte) (dur time.Duration) {
	lex := NewLexer(string(body))
	out := lex.Run()
	inTag := false
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
				dur += time.Duration(h*60*60) * time.Second
			}
		case Minutes:
			if !inTag {
				m, _ := strconv.Atoi(p.Val)
				dur += time.Duration(m*60) * time.Second
			}
		case Error:
			fmt.Println("Error", p.Val)
		}
	}
	return
}

type Tagged struct {
	Duration time.Duration
	Tag      string
}

func (tagged Tagged) String() string {
	dur := tagged.Duration
	hh := dur.Truncate(time.Hour)
	mm := dur - hh
	return fmt.Sprintf("%02v:%02v %s", hh.Hours(), mm.Minutes(), tagged.Tag)
}

func (par *Parser) SumTagged(body []byte) []Tagged {
	tagged := make([]Tagged, 0)
	tagDur := make(map[string]time.Duration, 0)
	lex := NewLexer(string(body))
	out := lex.Run()
	inTag := false
	var operator int = 1 // +1 or -1
	var dur time.Duration
	for {
		p, more := <-out
		if !more {
			break
		}
		switch p.Tok {
		case LeftParenthesis, RightParenthesis:
			inTag = !inTag
		case Operator:
			if p.Val == "-" {
				operator = -1
			}
		case Tag:
			if _, exists := tagDur[p.Val]; !exists {
				tagDur[p.Val] = 0
			}
			tagDur[p.Val] += dur
			dur = 0
			operator = 1
		case Hours:
			if inTag {
				h, _ := strconv.Atoi(p.Val)
				dur += time.Duration(h*60*60*operator) * time.Second
			}
		case Minutes:
			if inTag {
				m, _ := strconv.Atoi(p.Val)
				dur += time.Duration(m*60*operator) * time.Second
			}
		case Error:
			fmt.Println("Error", p.Val)
		}
	}
	for tag, dur := range tagDur {
		tagged = append(tagged, Tagged{dur, tag})
	}
	return tagged
}
