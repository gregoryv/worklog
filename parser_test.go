package timesheet

func ExampleParser_Dump() {
	par := NewParser()
	sheet := `2018 January
----------
1  1 Mon 8 (4 semester) thailand (2 pool)`
	par.Dump([]byte(sheet))
	//output:
	// Year[1,1]: "2018"
}
