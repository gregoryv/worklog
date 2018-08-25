package timesheet

func ExampleParser_Dump() {
	par := NewParser()
	sheet := `2018 January
----------
1  1 Mon 8 (8 semester)`
	par.Dump([]byte(sheet))
	//output:
	// Year[1,1]: "2018"
}
