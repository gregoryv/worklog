package main

import (
	"io"
	"text/template"
)

func renderText(w io.Writer, view *View, templatePath string) error {
	var t *template.Template
	var err error
	if templatePath != "" {
		t, err = template.ParseFiles(templatePath)
	} else {
		t = template.New("default")
		t, err = t.Parse(`{{range .Sheets}}{{.}}
{{end}}
Sum:           {{with .Reported}}{{.}}{{end}}
`)
	}
	if err != nil {
		return err
	}
	return t.Execute(w, view)
}
