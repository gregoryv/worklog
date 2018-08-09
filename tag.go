package timesheet

// tag enumerates recognized parts which the lexer may find
type tag int

//go:generate stringer -type tag tag.go
const (
	Unknown tag = iota
	Comment
	Text
	Year
	Month
	Separator
)
