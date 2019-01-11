package main

import (
	"fmt"
	"html/template"
	"io"
)

func renderHtml(w io.Writer, view *View, templatePath string) error {
	if templatePath == "" {
		return fmt.Errorf("Missing template")
	}
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	return t.Execute(w, view)
}
