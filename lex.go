package timesheet

type Lexer struct {
	scanner *Scanner
	C       chan Part
}

func NewLexer(txt string) *Lexer {
	return &Lexer{
		scanner: NewScanner(txt),
		C:       make(chan Part),
	}
}

func (l *Lexer) Run() chan Part {
	start := lexYear
	go l.run(start, l.scanner, l.C)
	return l.C
}

type lexFn func(s *Scanner) (p Part, next lexFn)

func (l *Lexer) run(start lexFn, s *Scanner, C chan Part) {
	for p, next := start(s); next != nil; p, next = next(s) {
		// Only report interesting tokens
		if p.Tok != Undefined {
			C <- p
		}
	}
	close(C)
}
