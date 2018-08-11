package timesheet

type Token int

//go:generate stringer -type Token token.go
const (
	Unknown Token = iota
	Number
)
