package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/abassGarane/snippet/pkg/forms"
	"github.com/abassGarane/snippet/pkg/models"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Println("NotFound page called")
		// http.NotFound(w, r)
    app.NotFound(w)
		return
	}
	app.infoLog.Println("Home snippet handler called...")
  // getting all snippets from database
  s, err := app.snippets.Latest()
  if err != nil{
    app.ServerError(w,err)
    return
  }
  app.render(w,r,"home.page.tmpl", templateData{
    Snippets: s,
  })
}

func (app *application) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Show snippet handler called...")
  id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		// http.NotFound(w, r)
    app.NotFound(w)
		return
	}
  s, err := app.snippets.Get(id)
  if err == models.ErrorNoRecord{
    app.NotFound(w)
    return
  }else if err != nil{
    app.ServerError(w,err)
    return
  }
  flash := app.session.PopString(r, "flash")
  app.render(w,r,"show.page.tmpl", templateData{
    Snippet: s,
    Flash: flash,
  })
}

func (app *application) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("CreateSnippet snippet handler called...")

  if err := r.ParseForm(); err != nil{
    app.ClientError(w,http.StatusBadRequest)
    return
  }
  form := forms.New(r.PostForm)
  form.Required("title", "content", "expires")
  form.MaxLength("title", 100)
  form.PermittedValues("expires", "365", "7", "1")

  // invalid form
  if !form.Valid(){
    app.render(w,r,"create.page.tmpl", templateData{Form: form})
    return
  }
  
  id , err := app.snippets.Insert(form.Get("title"),form.Get("content"), form.Get("expires"))
  if err != nil{
    app.ServerError(w,err)
    return
  }
  app.session.Put(r, "flash", "Snippet successfully created")
  http.Redirect(w,r,fmt.Sprintf("/snippet/%d",id), http.StatusSeeOther)
}

func (app *application)CreateSnippetForm(w http.ResponseWriter, r *http.Request)  {
  app.render(w,r,"create.page.tmpl", templateData{
    Form: forms.New(nil),
  })
}
