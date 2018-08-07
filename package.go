package timesheet

type lexer struct {
	report chan part
}

type state func(*lexer) state

type tag int

type part struct {
	tag  tag
	pos  int
	val  string
	line int
}

const (
	Unknown tag = iota
	Comment
	Text
)
