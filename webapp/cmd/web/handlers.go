package main

import (
	"html/template"
	"net/http"
	"path"
)

var (
	pathToTemlpates = "./templates/"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "home.page.gohtml", &TemplateData{})
}

type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {
	// parse the template from disk
	parsedTmpl, err := template.ParseFiles(path.Join(pathToTemlpates, t))
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return err
	}

	data.IP = app.ipFromContext(r.Context())

	// execute the template, passing it data if any
	err = parsedTmpl.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}
