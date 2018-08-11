package timesheet

func (l *Lexer) run(s *Scanner, out chan Part) {
	pos := s.Pos()
	val := s.ScanAll("0123456789")
	out <- Part{tok: Number, val: val, pos: pos}
	// todo define the grammar now...
}
