package timesheet

type lexer struct {
	report chan part
}

// state funcs are used to scan for parts from specific context
type state func(*lexer) state

// tag enumerates recognized parts which the lexer may find
type tag int

type part struct {
	tag tag
	val string
}

const (
	Unknown tag = iota
	Comment
	Text
)
