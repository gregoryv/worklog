package timesheet

type Logger interface {
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}

type NilLogger int

func (l NilLogger) Log(args ...interface{})                 {}
func (l NilLogger) Logf(format string, args ...interface{}) {}
