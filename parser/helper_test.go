// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package parser

import (
	"fmt"
)

func catchPanic(fn func()) (err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("%s", err)
		}
	}()
	fn()
	return
}
