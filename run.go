package timesheet

type lexFn func(s *Scanner, out chan Part) lexFn

func lexMonth(s *Scanner, out chan Part) lexFn {
	pos := s.Pos()
	// todo err check here
	first, _ := s.Scan("JFMASOND")
	rest, _ := s.ScanAll("abcdefghijklmnopqrstuvxyz")
	out <- Part{Tok: Month, Val: first + rest, Pos: pos}
	return nil
}

func lexYear(s *Scanner, out chan Part) lexFn {
	pos := s.Pos()
	val, ok := s.ScanAll("0123456789")
	p := Part{Tok: Number, Val: val, Pos: pos}
	if !ok {
		p.Tok = Error
		p.Val = "invalid year"
	}
	out <- p
	s.Scan(" ")
	return nil
}

func (l *Lexer) run(s *Scanner, out chan Part) {
	// We expect to start the file with a year
	for fn := lexYear; fn != nil; fn = fn(s, out) {
	}
	close(out)
}
