package main

import "net/http"
func (app *application)Routes()http.Handler  {

  mux := http.NewServeMux()

  mux.HandleFunc("/", app.Home)
  mux.HandleFunc("/snippet",app.ShowSnippet)
  mux.HandleFunc("/snippet/create",app.CreateSnippet)

  fileServer := http.FileServer(http.Dir("./ui/static"))
  mux.Handle("/static/", http.StripPrefix("/static",fileServer))

  return app.logRequest(secureHeaders(mux))
}

