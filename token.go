package timesheet

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
