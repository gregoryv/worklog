package main

import (
	"io"
	"text/template"

	timesheet "github.com/gregoryv/go-timesheet"
)

func renderText(w io.Writer, report *timesheet.Report, templatePath string) error {
	var t *template.Template
	var err error
	if templatePath != "" {
		t, err = template.ParseFiles(templatePath)
	} else {
		t = template.New("default")
		t, err = t.Parse(`{{range .Sheets}}{{.}}
{{end}}
{{with .SumReported}}{{.}}{{end}}
`)
	}
	if err != nil {
		return err
	}
	return t.Execute(w, report)
}
