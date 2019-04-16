// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package main

import (
	"fmt"
	"time"

	timesheet "github.com/gregoryv/go-timesheet"
)

func Example_ConvertToTagView() {
	tag := timesheet.Tagged{60 * time.Second, "vacation"}
	in := []timesheet.Tagged{tag}
	out := ConvertToTagView(in)
	fmt.Println(out[0].Duration, out[0].Tag)
	//output: 0:01 vacation
}
