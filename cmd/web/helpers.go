package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application)ServerError(w http.ResponseWriter, err error)  {
  trace := fmt.Sprintf("%s\n%s",err.Error(), debug.Stack())
  app.errorLog.Output(2,trace) 
  http.Error(w,http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) ClientError(w http.ResponseWriter, status int)  {
  http.Error(w,http.StatusText(status),status)
}

func (app *application)NotFound(w http.ResponseWriter)  {
  app.ClientError(w,http.StatusNotFound)
}

func (app *application)render(w http.ResponseWriter, r *http.Request, name string, td templateData)  {
  ts, ok := app.templateCache[name]
  if !ok{
    app.ServerError(w,fmt.Errorf("Template %s does not exist", name))
    return
  }
  if err := ts.Execute(w, td); err != nil{
    app.ServerError(w,err)
  }
}
