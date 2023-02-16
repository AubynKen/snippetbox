package main

import (
	"github.com/AubynKen/snippetbox/internal/models"
	"html/template"
	"path/filepath"
	"time"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

func newTemplateData() *templateData {
	return &templateData{
		CurrentYear: time.Now().Year(),
	}
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template, 0)

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// Parse the base template file into a template set with registered
		// functions
		ts, err := template.New(name).
			Funcs(functions).
			ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}

		// Call ParseGlob() *on this template set* to add any partials.
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		// Call ParseFiles() *on this template set* to add the  page template.
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add the template set to the map as normal...
		cache[name] = ts
	}

	return cache, nil
}
