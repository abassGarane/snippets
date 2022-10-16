package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/abassGarane/snippet/pkg/models"
	"github.com/justinas/nosurf"
)

func secureHeaders(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { 

    // secure http headers
    w.Header().Set("X-XSS-Protection", "1;mode=block")
    w.Header().Set("X-Frame-Options", "deny")

    next.ServeHTTP(w,r)
  })
}

func (app *application)logRequest(next http.Handler)http.Handler  {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { 
    app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL)

    next.ServeHTTP(w,r)
  })
}

func (app *application)recoverPanic(next http.Handler)http.Handler  {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { 

    defer func() { 
      if err := recover(); err != nil{
        w.Header().Set("Connection", "close")
        app.ServerError(w,fmt.Errorf("%s", err))
      }
    }()

    next.ServeHTTP(w,r)
  })
}

func (app *application)requireAuthentication(next http.Handler)http.Handler  {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  {
    if app.AuthenticatedUser(r) == nil{
      http.Redirect(w,r,"/user/login", 302)
      return
    }
    next.ServeHTTP(w,r)
  })
}

func (app *application)noSurf(next http.Handler)http.Handler  {
  csrfHandler := nosurf.New(next)
  csrfHandler.SetBaseCookie(http.Cookie{
    HttpOnly: true,
    Path: "/",
    Secure: true,
  })
  return csrfHandler
}
func (app *application)authenticate(next http.Handler)http.Handler  {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  {
    exists := app.session.Exists(r, "userID")
    if  !exists{
      next.ServeHTTP(w,r)
      return
    }

    user, err := app.users.Get(app.session.GetInt(r, "userID"))
    if err == models.ErrorNoRecord{
      app.session.Remove(r, "userID")
      return
    }else if err != nil{
      app.ServerError(w,err)
      return
    }
    ctx := context.WithValue(r.Context(), contextKeyUser, user) 
    next.ServeHTTP(w,r.WithContext(ctx))
  })
}
