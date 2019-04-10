// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.
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
