package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Println("NotFound page called")
		// http.NotFound(w, r)
    app.NotFound(w)
		return
	}
	app.infoLog.Println("Home snippet handler called...")
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
    // app.errorLog.Println(err.Error())
		// http.Error(w, "Internal server error..", http.StatusInternalServerError)
    app.ServerError(w,err)
    return
	}
	if err = ts.Execute(w, nil); err != nil {
		// app.errorLog.Println(err.Error())
		// http.Error(w, "Internal server error...", http.StatusInternalServerError)
    app.ServerError(w,err) 
    return
	}
}

func (app *application) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Show snippet handler called...")
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		// http.NotFound(w, r)
    app.NotFound(w)
		return
	}
	// w.Write([]byte("Display a specific snippet ..."))
	fmt.Fprintf(w, "Display a specific snippet with id %d...", id)
}

func (app *application) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("CreateSnippet snippet handler called...")
	// Removing a header
	w.Header()["Date"] = nil
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Method not allowed"))
		// http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
    app.ClientError(w,http.StatusMethodNotAllowed)
		return
	}
	// Update or modify an existing header
	// w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Create a new snippet..."))
}