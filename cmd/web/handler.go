package main

import (
	"net/http"

	"backen/internals"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	users := internals.NewUsers()

	w.Header().Set("Content-Type", "text/html")
	
	data := make(map[string]interface{})
	data["datas"] = users.AllUsers()

	if err := app.render(w, r, "home", PageData{
		Title: "meine Hausseite",
		Data:  data,
	}); err != nil {
		app.Err.Println(err)
	}

}

func (app *application) about(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	if err := app.render(w, r, "about", PageData{
		Title: "About Seite",
	}); err != nil {
		app.Err.Println(err)
	}

}
