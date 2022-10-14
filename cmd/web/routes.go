package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)
func (app *application)Routes()http.Handler  {

  // middlewares
  standardMiddlewares := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

  // Registering a router
  mux := pat.New()
  // mux := http.NewServeMux()

  mux.Get("/", http.HandlerFunc(app.Home))
  mux.Get("/snippet/create", http.HandlerFunc(app.CreateSnippetForm))
  mux.Post("/snippet/create",http.HandlerFunc(app.CreateSnippet))
  mux.Get("/snippet/:id",http.HandlerFunc(app.ShowSnippet))

  fileServer := http.FileServer(http.Dir("./ui/static"))
  mux.Get("/static/", http.StripPrefix("/static",fileServer))

  return standardMiddlewares.Then(mux)
}

