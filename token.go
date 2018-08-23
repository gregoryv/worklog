package timesheet

type Token int

//go:generate stringer -type Token token.go
const (
	Undefined Token = iota
	Error
	Number
	Month
	Separator
	Day
	Hour
	LeftParenthesis
	Operator // -,+
)
