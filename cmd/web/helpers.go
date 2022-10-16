package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
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

func (app *application)addDefaultData(td *templateData, r *http.Request)*templateData  {
  if td == nil{
    td = &templateData{}
  }
  td.CurrentYear = time.Now().Year() 
  td.Flash = app.session.PopString(r, "flash")
  td.AuthenticatedUser = app.AuthenticatedUser(r)
  return td
}

func (app *application)render(w http.ResponseWriter, r *http.Request, name string, td templateData)  {
  ts, ok := app.templateCache[name]
  if !ok{
    app.ServerError(w,fmt.Errorf("Template %s does not exist", name))
    return
  }
  buf := new(bytes.Buffer)
  if err := ts.Execute(buf, app.addDefaultData(&td,r)); err != nil{
    app.ServerError(w,err)
    return
  }
  buf.WriteTo(w)
  
}

func (app *application)AuthenticatedUser(r *http.Request)int  {
  return app.session.GetInt(r, "userID")
}
