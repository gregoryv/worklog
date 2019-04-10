// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.
package parser

func ExampleParser_Dump() {
	NewParser().Dump([]byte(`2018 January
----------
1  1 Mon 8:10 (4 semester) was in thailand (2 pool)
`))
	// output:
	// Year[1,1]: "2018"
	// Month[1,6]: "January"
	// Separator[2,1]: "----------"
	// Week[3,1]: "1"
	// Date[3,4]: "1"
	// Day[3,6]: "Mon"
	// Hours[3,10]: "8"
	// Colon[3,11]: ":"
	// Minutes[3,12]: "10"
	// LeftParenthesis[3,15]: "("
	// Hours[3,16]: "4"
	// Tag[3,18]: "semester"
	// RightParenthesis[3,26]: ")"
	// Note[3,28]: "was in thailand "
	// LeftParenthesis[3,44]: "("
	// Hours[3,45]: "2"
	// Tag[3,47]: "pool"
	// RightParenthesis[3,51]: ")"
}

func ExampleParser_Dump_bad() {
	NewParser().Dump([]byte(`2018 nosuchmonth`))
	// output:
	// Year[1,1]: "2018"
	// Error[1,6]: "invalid Month"
	// Error[1,6]: "invalid Separator"
	// Error[1,6]: "invalid Week"
	// Error[1,6]: "invalid Date"
	// Error[1,6]: "invalid Day"
	// Note[1,6]: "nosuchmonth"
}
