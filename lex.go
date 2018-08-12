package timesheet

type Lexer struct {
	scanner *Scanner
	out     chan Part
}

func (l *Lexer) Run() chan Part {
	go l.run(lexYear, l.scanner, l.out)
	return l.out
}

func NewLexer(txt string) *Lexer {
	return &Lexer{
		scanner: NewScanner(txt),
		out:     make(chan Part),
	}
}
