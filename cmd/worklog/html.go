// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.
package main

import (
	"fmt"
	"html/template"
	"io"
)

func renderHtml(w io.Writer, view *ReportView, templatePath string) error {
	if templatePath == "" {
		return fmt.Errorf("Missing template")
	}
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	return t.Execute(w, view)
}
