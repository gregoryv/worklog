package timesheet

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

type Sheet struct {
	Period   string // Year Month
	Reported Tagged
	Tags     []Tagged
}

func (par *Parser) Parse(body []byte) (sheet *Sheet, err error) {
	sheet = &Sheet{}
	lex := NewLexer(string(body))
	out := lex.Run()
	tagDur := make(map[string]time.Duration, 0)
	var dur time.Duration // for tags
	operator := 1         // +1 or -1
	tagged := make([]Tagged, 0)
	inTag := false
	for {
		p, more := <-out
		switch p.Tok {
		case LeftParenthesis, RightParenthesis:
			inTag = !inTag
		case Year:
			sheet.Period += p.Val
		case Month:
			sheet.Period += " " + p.Val
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
			h, _ := strconv.Atoi(p.Val)
			hh := time.Duration(h*60*60*operator) * time.Second
			if inTag {
				dur += hh
			} else {
				sheet.Reported.Duration += hh
			}
		case Minutes:
			m, _ := strconv.Atoi(p.Val)
			mm := time.Duration(m*60*operator) * time.Second
			if inTag {
				dur += mm
			} else {
				sheet.Reported.Duration += mm
			}
		case Error:
			err = fmt.Errorf("%s", p)
		}
		if !more || err != nil {
			break
		}
	}
	for tag, dur := range tagDur {
		tagged = append(tagged, Tagged{dur, tag})
	}
	sort.Sort(byTag(tagged))
	sheet.Tags = tagged
	return
}
