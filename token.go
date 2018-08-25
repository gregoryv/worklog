package timesheet

type Token int

//go:generate stringer -type Token token.go
const (
	Undefined Token = iota
	Error
	Number
	Hours
	Month
	Separator
	Day
	Hour
	LeftParenthesis
	Operator // -,+
)
