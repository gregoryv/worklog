package timesheet

type Pos struct {
	line, column int
}

func NewPos() *Pos {
	return &Pos{}
}
