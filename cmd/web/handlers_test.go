package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T)()  {
  rr := httptest.NewRecorder()
  r,err := http.NewRequest("GET", "/", nil)
  if err !=nil{
    t.Fatal(err)
  }
  Ping(rr,r)

  rs := rr.Result() 
  if rs.StatusCode != http.StatusOK{
    t.Errorf("Want %d; got %d", http.StatusOK, rs.StatusCode)
  }

  defer rs.Body.Close()
  body, err := ioutil.ReadAll(rs.Body)
  if err != nil{
    t.Fatal(err)
  }

  if string(body) != "ok"{
    t.Errorf("Want the body to equal %q", "ok")
  }
  
}
