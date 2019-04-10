// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.
package token

type Token int

//go:generate stringer -type Token token.go
const (
	Undefined Token = iota
	Error
	Year
	Hours
	Note
	Month
	Separator
	Day
	Date
	Hour
	LeftParenthesis
	RightParenthesis
	Operator // -,+
	Colon
	Minutes
	Tag
	Week
)
