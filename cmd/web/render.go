package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

//go:embed templates
var templatesFS embed.FS

var functions = template.FuncMap{
	"Join":      ConcatString,
	"ToString":  ToString,
	"YearMonth": CurrentYear,
}

func readFiles() []string {
	f, err := filepath.Glob("cmd/web/templates/**/*.html")
	var files = []string{
		"templates/index.html",
	}
	for _, fl := range f {
		files = append(files, strings.Replace(fl, "cmd/web/", "", -1))
	}
	if err != nil {
		fmt.Println(err)
	}
	return files
}

func ConcatString(s ...string) string {
	return strings.Trim(strings.Join(s, ""), " ")
}

func ToString(s any) string {
	return fmt.Sprintf("%v", s)
}

func CurrentYear() string {
	return fmt.Sprintf("%v/%v", time.Now().UTC().Month(), time.Now().UTC().Year())
}

type PageData struct {
	Title string
	Data  map[string]interface{}
}

func (app *application) render(w http.ResponseWriter, r *http.Request, page string, td PageData) error {
	var tmpl *template.Template
	var err error
	partials := readFiles()	
	tmpl, err = template.New(fmt.Sprintf("%s.html", page)).Funcs(functions).ParseFS(templatesFS, partials...)
	if err != nil {
		return err
	}

	if err = tmpl.Execute(w, td); err != nil {
		return err
	}
	return nil
}
