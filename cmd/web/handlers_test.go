package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestPing(t *testing.T)()  {
  app := NewTestApplication(t)
  ts := NewTestServer(t, app.Routes())
  defer ts.Close()
  code, _, body := ts.Get(t, "/ping") 
  // t.Fatal(ts.URL)
  if code != http.StatusOK {
    t.Errorf("want %d; got %d", http.StatusOK, code)
  }

  if string(body) != "ok"{
    t.Errorf("want the body to be equal to %q", "ok")
  }
  
}

func TestShowSnippet(t *testing.T)  {
  app := NewTestApplication(t)
  ts := NewTestServer(t, app.Routes())

  defer ts.Close() 

  tests := []struct{
    name string
    urlPath string
    wantCode int
    wantBody []byte
  }{
    {"Valid ID", "/snippet/1", http.StatusOK, []byte("A shattered visage lies...")},
    {"Non-Existing ID", "/snippet/2", http.StatusNotFound, nil},
    {"Decimal ID", "/snippet/2.55", http.StatusNotFound, nil},
    {"Empty ID", "/snippet/", http.StatusNotFound, nil},
  }
  for _, tt := range tests{
    t.Run(tt.name, func(t *testing.T) { 
      code, _, body := ts.Get(t,tt.urlPath)
      if code != tt.wantCode{
        t.Errorf("want %d; got %d", tt.wantCode, code)
      }
      if !bytes.Contains(body, tt.wantBody){
        t.Errorf("want body to contain %q", tt.wantBody)
      }
    })
  }
}
