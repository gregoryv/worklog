package timesheet

import (
	"testing"
)

func Test_lexYear(t *testing.T) {
	cases := []struct {
		s   *Scanner
		out chan Part
	}{
		{NewScanner("2018"), make(chan Part)},
	}
	for _, c := range cases {
		go lexYear(c.s, c.out)
		part := <-c.out
		assert(t, "",
			equals("", "Number[1,1]: \"2018\"", part.String()),
		)
		close(c.out)
	}
}
