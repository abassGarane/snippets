package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/abassGarane/snippet/pkg/forms"
	"github.com/abassGarane/snippet/pkg/models"
)

func Ping(w http.ResponseWriter, r *http.Request)  {
  w.Write([]byte("ok"))
}

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
  app.render(w,r,"show.page.tmpl", templateData{
    Snippet: s,
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

func (app *application)SignupUserForm(w http.ResponseWriter, r *http.Request)  {
  app.render(w,r,"signup.page.tmpl", templateData{ 
    Form: forms.New(nil),
  })
}

func (app *application)SignupUser(w http.ResponseWriter, r *http.Request)  {
  err := r.ParseForm() 
  if err != nil{
    app.ClientError(w,http.StatusBadRequest)
    return
  }
  form := forms.New(r.PostForm)
  form.Required("name", "email", "password")
  form.MinLength("password", 10)
  form.MatchesPattern("email", forms.EmailXR)

  if !form.Valid(){
    app.render(w,r,"signup.page.tmpl", templateData{Form: form})
    return
  }
  err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"))
  if err == models.ErrorDuplicateEmail{
    form.Errors.Add("email", "Address already in use")
    app.render(w,r,"signup.page.tmpl",templateData{
      Form: form,
    })
    return
  }else if err != nil{
    app.ServerError(w,err)
    return
  }
  // user successfuly saved in db
  app.session.Put(r,"flash", "Your signup was successful. please log in")
  http.Redirect(w,r,"/user/login",http.StatusSeeOther)
}

func (app *application)LoginUserForm(w http.ResponseWriter, r *http.Request)  {
  app.render(w,r,"login.page.tmpl", templateData{ 
    Form: forms.New(nil),
  })
}

func (app *application)LoginUser(w http.ResponseWriter, r *http.Request)  {
  if err := r.ParseForm(); err != nil{
    app.ClientError(w,http.StatusBadRequest)
    return
  }
  form := forms.New(r.PostForm)
  id, err := app.users.Authenticate(form.Get("email"), form.Get("password"))
  if err == models.ErrorInvalidCredentials{
    form.Errors.Add("generic", "Email or Password is incorrect")
    app.render(w,r,"login.page.tmpl", templateData{ 
      Form: form,
    })
    return
  }else if err != nil{
    app.ServerError(w,err)
    return
  }
  app.session.Put(r, "userID", id)
  http.Redirect(w,r,"/",http.StatusSeeOther )
}

func (app *application)LogoutUser(w http.ResponseWriter, r *http.Request)  {
  app.session.Remove(r, "userID")
  app.session.Put(r, "flash", "You have been logged out successfuly")
  fmt.Fprintln(w,"Logout user...")
  http.Redirect(w,r,"/", 303)
}
