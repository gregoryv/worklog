// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.
package timesheet

import (
	"bytes"
	"flag"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/gregoryv/find"
)

var addCopy = flag.Bool("add-copy", false, "Prepends copy statement, see license_test.go")

func Test_check_license_in_go_files(t *testing.T) {
	result, _ := find.ByName("*.go", ".")
	missing := make([]string, 0)
	pattern := `// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.`
	for e := result.Front(); e != nil; e = e.Next() {
		if file, ok := e.Value.(string); ok {
			body, err := ioutil.ReadFile(file)
			if err != nil {
				t.Fatal(err)
			}
			if bytes.Index(body, []byte(pattern)) == -1 {
				if *addCopy {
					tmp := append([]byte(pattern), '\n')
					body = append(tmp, body...)
					ioutil.WriteFile(file, body, 0622)
				} else {
					missing = append(missing, file)
				}
			}
		}
	}
	if len(missing) > 0 {
		t.Errorf("Missing Copy in \n%s\n\nfix with -add-copy", strings.Join(missing, "\n"))
	}
}
