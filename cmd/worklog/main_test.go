// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.
package main

import (
	"bytes"
	"os/exec"
	"testing"

	"github.com/gregoryv/asserter"
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
