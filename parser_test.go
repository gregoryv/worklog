package timesheet

import (
	"fmt"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	for i, src := range okTimesheets {
		p := NewParser()
		p.Debug = t
		if err := p.Parse(strings.NewReader(src), fmt.Sprintf("case %v", i)); err != nil {
			t.Error(err)
		}
		t.Fail()
	}
}

var okTimesheets = []string{`
2018 August
`,
}
