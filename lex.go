package timesheet

type Lexer struct {
	name    string // eg. named file, todo remove from here
	scanner *Scanner
	out     chan Part
}

func (l *Lexer) Run() chan Part {
	go l.run(l.scanner, l.out)
	return l.out
}

func NewLexer(name, txt string) *Lexer {
	return &Lexer{
		name:    name,
		scanner: NewScanner(txt),
		out:     make(chan Part),
	}
}
