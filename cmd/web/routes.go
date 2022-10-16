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

  mux.Get("/", dynamicMiddleware.ThenFunc(app.Home))
  mux.Get("/snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.CreateSnippetForm))
  mux.Post("/snippet/create",dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.CreateSnippet))
  mux.Get("/snippet/:id",dynamicMiddleware.ThenFunc(app.ShowSnippet))

  // user Routes
  mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.SignupUserForm))
  mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.SignupUser))
  mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.LoginUserForm))
  mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.LoginUser))
  mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.LogoutUser))

  fileServer := http.FileServer(http.Dir("./ui/static"))
  mux.Get("/static/", http.StripPrefix("/static",fileServer))

  return standardMiddlewares.Then(mux)
}

