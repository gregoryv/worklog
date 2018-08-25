package timesheet

type Token int

//go:generate stringer -type Token token.go
const (
	Undefined Token = iota
	Error
	Number
	Year
	Hours
	Note
	Month
	Separator
	Day
	Hour
	LeftParenthesis
	RightParenthesis
	Operator // -,+
	Colon
	Minutes
)
