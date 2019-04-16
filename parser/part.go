// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package parser

import (
	"fmt"

	"github.com/gregoryv/go-timesheet/token"
)

type Part struct {
	Tok token.Token
	Val string
	Pos Position
}

func (p *Part) Defined() bool {
	return p.Tok != token.Undefined
}

func (a Part) Equals(b Part) bool {
	return a.Tok == b.Tok &&
		a.Val == b.Val &&
		a.Pos.Equals(b.Pos)
}

func (p *Part) Errorf(format string, args ...interface{}) error {
	p.Val = fmt.Sprintf(format, args...)
	p.Tok = token.Error
	return fmt.Errorf(p.Val)
}

func (p Part) String() string {
	return fmt.Sprintf("%s[%s]: %q", p.Tok, p.Pos.String(), p.Val)
}
