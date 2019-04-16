// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package main

import (
	"io"
	"text/template"
)

func renderText(w io.Writer, view *ReportView, templatePath string) error {
	var t *template.Template
	var err error
	if templatePath != "" {
		t, err = template.ParseFiles(templatePath)
	} else {
		t = template.New("default")
		t, err = t.Parse(`{{range .Sheets}}{{.Period}} {{.Reported}} {{.Diff}} {{range .Tags}} ({{.}}){{end}}
{{end}}
{{printf "%22s" .ReportedIndent}} {{.Diff}}
{{range .Tags}}{{printf "%30s" ""}} {{.Duration}} {{.Tag}}
{{end}}`)
	}
	if err != nil {
		return err
	}
	return t.Execute(w, view)
}
