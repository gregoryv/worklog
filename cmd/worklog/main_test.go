// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
	"time"

	"github.com/gregoryv/asserter"
	timesheet "github.com/gregoryv/go-timesheet"
	"github.com/gregoryv/golden"
)

func TestFeature(t *testing.T) {
	out, err := exec.Command("worklog", "-origin", "../../assets/orig2018",
		"../../assets/201801.timesheet").CombinedOutput()
	assert := asserter.New(t)
	assert(err == nil).Fatal(err, string(out))

	period := "2018 January"
	year := bytes.Index(out, []byte(period))
	assert(year == 0).Errorf("Did not start with %q", period)
}

func Example_ConvertToTagView() {
	tag := timesheet.Tagged{60 * time.Second, "vacation"}
	in := []timesheet.Tagged{tag}
	out := ConvertToTagView(in)
	fmt.Println(out[0].Duration, out[0].Tag)
	//output: 0:01 vacation
}

func Test_writeText(t *testing.T) {
	w := bytes.NewBufferString("")
	writeText(w, "", "../../testdata/orig", []string{"../../testdata/201506.timesheet"})
	golden.Assert(t, w.String())
}
