package timesheet

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Sheet struct {
	Period   string // Year Month
	Reported Tagged
	Tags     []Tagged
}

func NewSheet() *Sheet {
	return &Sheet{Reported: Tagged{0, "reported"}}
}

func (par *Parser) Parse(body []byte) (sheet *Sheet, err error) {
	sheet = NewSheet()
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
			hh := time.Duration(h*operator) * time.Hour
			if inTag {
				dur += hh
			} else {
				sheet.Reported.Duration += hh
			}
		case Minutes:
			m, _ := strconv.Atoi(p.Val)
			mm := time.Duration(m*operator) * time.Minute
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

func (sheet *Sheet) String() string {
	return fmt.Sprintf("%s %s %s", sheet.Period, sheet.Reported,
		strings.Join(inParenthesis(sheet.Tags), " "))
}

func inParenthesis(tags []Tagged) []string {
	res := make([]string, 0)
	for _, tag := range tags {
		res = append(res, fmt.Sprintf("(%s)", tag))
	}
	return res
}
