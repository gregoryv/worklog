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
		t, err = t.Parse(`{{range .Sheets}}{{.Period}} {{.Reported}} {{.Diff}}{{range .Tags}} ({{.}}){{end}}
{{end}}
{{.Reported}} {{.Diff}}
`)
	}
	if err != nil {
		return err
	}
	return t.Execute(w, view)
}
