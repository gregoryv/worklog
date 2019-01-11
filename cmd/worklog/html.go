package main

import (
	"fmt"
	"html/template"
	"io"

	timesheet "github.com/gregoryv/go-timesheet"
)

func renderHtml(w io.Writer, report *timesheet.Report, templatePath string) error {
	if templatePath == "" {
		return fmt.Errorf("Missing template")
	}
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	return t.Execute(w, report)
}
