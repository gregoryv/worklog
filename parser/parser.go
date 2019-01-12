package parser

import (
	"fmt"
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
