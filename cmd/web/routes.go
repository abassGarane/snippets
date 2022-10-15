package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)
func (app *application)Routes()http.Handler  {

  // middlewares
  standardMiddlewares := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
  dynamicMiddleware := alice.New(app.session.Enable)

  // Registering a router
  mux := pat.New()
  // mux := http.NewServeMux()

  mux.Get("/", dynamicMiddleware.ThenFunc(app.Home))
  mux.Get("/snippet/create", dynamicMiddleware.ThenFunc(app.CreateSnippetForm))
  mux.Post("/snippet/create",dynamicMiddleware.ThenFunc(app.CreateSnippet))
  mux.Get("/snippet/:id",dynamicMiddleware.ThenFunc(app.ShowSnippet))

  fileServer := http.FileServer(http.Dir("./ui/static"))
  mux.Get("/static/", http.StripPrefix("/static",fileServer))

  return standardMiddlewares.Then(mux)
}

