package timesheet

type Token int

//go:generate stringer -type Token token.go
const (
	Error Token = iota
	Number
	Month
	Separator
)
