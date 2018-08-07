package timesheet

import (
	"fmt"
)

func start(l *lexer) state {
	l.report <- part{tag: Comment, val: "start"}
	return mid
}

func mid(l *lexer) state {
	l.report <- part{tag: Text, val: "mid"}
	return end
}

func end(l *lexer) state {
	l.report <- part{tag: Text, val: "end"}
	return nil
}

func Example() {
	l := &lexer{report: make(chan part)}
	go func() {
		for {
			select {
			case p := <-l.report:
				fmt.Println(p)
			}
		}
	}()
	for state := start; state != nil; {
		state = state(l)
	}
	//output:
	// start
	// mid
	// end
}
