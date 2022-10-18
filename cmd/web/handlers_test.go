package main

import (
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
