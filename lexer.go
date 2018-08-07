package timesheet

import (
	"fmt"
)

func (p part) String() string {
	return fmt.Sprintf("%v", p.val)
}
