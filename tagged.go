package timesheet

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

type Tagged struct {
	Duration time.Duration
	Tag      string
}

func (tagged Tagged) String() string {
	dur := tagged.Duration
	hh := dur.Truncate(time.Hour)
	var operator time.Duration = 1
	if hh < 0 {
		operator = -1
	}
	mm := (dur - hh) * operator
	return fmt.Sprintf("%v:%02v %s", hh.Hours(), mm.Minutes(), tagged.Tag)
}

func (par *Parser) SumTagged(body []byte) ([]Tagged, error) {
	tagged := make([]Tagged, 0)
	tagDur := make(map[string]time.Duration, 0)
	lex := NewLexer(string(body))
	out := lex.Run()
	inTag := false
	var operator int = 1 // +1 or -1
	var dur time.Duration
	for {
		p, more := <-out
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
			return tagged, fmt.Errorf("%s", p)
		}
		if !more {
			break
		}
	}
	for tag, dur := range tagDur {
		tagged = append(tagged, Tagged{dur, tag})
	}
	sort.Sort(byTag(tagged))
	return tagged, nil
}

type byTag []Tagged

func (by byTag) Len() int           { return len(by) }
func (by byTag) Less(i, j int) bool { return by[i].Tag < by[j].Tag }
func (by byTag) Swap(i, j int) {
	by[i], by[j] = by[j], by[i]
}
