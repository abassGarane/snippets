package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T)()  {
  app := &application{ 
    errorLog: log.New(io.Discard, "", 0),
    infoLog: log.New(io.Discard, "", 0),
  } 

  ts := httptest.NewTLSServer(app.Routes())

  defer ts.Close()
  rs, err := ts.Client().Get(ts.URL + "/ping")
  if err != nil{
    t.Fatal(err)
  }
  // t.Fatal(ts.URL)
  if rs.StatusCode != http.StatusOK {
    t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
  }

  defer rs.Body.Close() 
  body, err := ioutil.ReadAll(rs.Body)
  if err != nil{
    t.Fatal(err)
  }
  if string(body) != "ok"{
    t.Errorf("want the body to be equal to %q", "ok")
  }
  
}
