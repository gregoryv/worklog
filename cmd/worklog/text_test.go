// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package main

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/golden"
)

func Test_writeText(t *testing.T) {
	w := bytes.NewBufferString("")
	writeText(w, "", "../../testdata/orig", []string{"../../testdata/201506.timesheet"})
	got := w.String()
	exp := golden.LoadString()
	if got != exp {
		assert := asserter.New(t)
		assert().Equals(got, exp)
	}
	golden.SaveString(t, got)
}
