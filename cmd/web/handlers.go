package main

import (
	"fmt"
	// "html/template"
	"log"
	"net/http"
	"strconv"

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
  // for _, snippet := range s{
  //   fmt.Fprintf(w, "%v\n", snippet)
  // }
 //  data := &templateData{ 
 //    Snippets: s,
 //  }

	// files := []string{
	// 	"./ui/html/home.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
 //    // app.errorLog.Println(err.Error())
	// 	// http.Error(w, "Internal server error..", http.StatusInternalServerError)
 //    app.ServerError(w,err)
 //    return
	// }
	// if err = ts.Execute(w, data); err != nil {
	// 	// app.errorLog.Println(err.Error())
	// 	// http.Error(w, "Internal server error...", http.StatusInternalServerError)
 //    app.ServerError(w,err) 
 //    return
	// }
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
 //  data := &templateData{Snippet: s}
 //  files := []string{
	// 	"./ui/html/show.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
 //    // app.errorLog.Println(err.Error())
	// 	// http.Error(w, "Internal server error..", http.StatusInternalServerError)
 //    app.ServerError(w,err)
 //    return
	// }
	// if err = ts.Execute(w, data); err != nil {
	// 	// app.errorLog.Println(err.Error())
	// 	// http.Error(w, "Internal server error...", http.StatusInternalServerError)
 //    app.ServerError(w,err) 
 //    return
	// }
	// w.Write([]byte("Display a specific snippet ..."))
	// fmt.Fprintf(w, "%v", s)
  app.render(w,r,"show.page.tmpl", templateData{
    Snippet: s,
  })
}

func (app *application) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("CreateSnippet snippet handler called...")
	// Removing a header
	w.Header()["Date"] = nil
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Method not allowed"))
		// http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
    app.ClientError(w,http.StatusMethodNotAllowed)
		return
	}
  title := "O snail"
  content := "O snail\nClimb Mount Kenya,\n But slowly, slowly!\n"
  expires := "7"
  id , err := app.snippets.Insert(title, content, expires)
  app.infoLog.Println(id)
  if err != nil{
    app.ServerError(w,err)
    return
  }
  http.Redirect(w,r,fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
	// Update or modify an existing header
	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte("Create a new snippet..."))
}

func (app *application)CreateSnippetForm(w http.ResponseWriter, r *http.Request)  {
  w.Write([]byte("Create a new snippet"))
}
