package timesheet

import (
	"fmt"
	"testing"
)

func todo(t *testing.T) {
	t.Helper()
	t.Errorf("TODO")
}

func catchPanic(fn func()) (err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("%s", err)
		}
	}()
	fn()
	return
}

func assert(t *testing.T, msg string, errors ...error) {
	t.Helper()
	for _, err := range errors {
		if err != nil {
			t.Errorf("%s: %s", msg, err)
		}
	}
}

func equals(label string, a, b interface{}) (err error) {
	if a != b {
		return fmt.Errorf("expected %s=\"%v\", got \"%v\"", label, a, b)
	}
	return
}
