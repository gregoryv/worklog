package timesheet

type lexFn func(s *Scanner, out chan Part) lexFn

func lexYear(s *Scanner, out chan Part) lexFn {
	pos := s.Pos()
	val := s.ScanAll("0123456789")
	out <- Part{tok: Number, val: val, pos: pos}
	s.Scan(" ")
	return nil
}

func (l *Lexer) run(s *Scanner, out chan Part) {
	// We expect to start the file with a year
	for fn := lexYear; fn != nil; fn = fn(s, out) {
	}
	close(out)
}
