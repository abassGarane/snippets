package main

import (
  "net/http"
  "github.com/justinas/alice"
)
func (app *application)Routes()http.Handler  {

  // middlewares
  standardMiddlewares := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

  mux := http.NewServeMux()

  mux.HandleFunc("/", app.Home)
  mux.HandleFunc("/snippet",app.ShowSnippet)
  mux.HandleFunc("/snippet/create",app.CreateSnippet)

  fileServer := http.FileServer(http.Dir("./ui/static"))
  mux.Handle("/static/", http.StripPrefix("/static",fileServer))

  return standardMiddlewares.Then(mux)
}

