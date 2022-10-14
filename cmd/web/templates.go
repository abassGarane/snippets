package main

import (
	"html/template"
	"path/filepath"

	"time"

	"github.com/abassGarane/snippet/pkg/forms"
	"github.com/abassGarane/snippet/pkg/models"
)

type templateData struct{
  Snippet *models.Snippet
  Snippets []*models.Snippet
  CurrentYear int
  // FormErrors map[string]string
  // FormData url.Values
  Form *forms.Form
}

func humanDate(t time.Time) string  {
  return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
  "humanDate":humanDate,
}

func newTemplateCache( dir string)(map[string]*template.Template, error)  {
  var cache = map[string]*template.Template{}

  pages , err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
  if err != nil{
    return nil, err
  }
  for _, page := range pages{
    name := filepath.Base(page)
    ts, err := template.New(name).Funcs(functions).ParseFiles(page)
    if err != nil{
      return nil, err
    }
    ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
    if err != nil{
      return nil, err
    }
    ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
    if err != nil{
      return nil, err
    }
    cache[name]=ts
  }
  return cache, nil
}
