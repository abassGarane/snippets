package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewTestApplication(t *testing.T)*application  {
 return &application{ 
    errorLog: log.New(io.Discard, "", 0),
    infoLog: log.New(io.Discard, "", 0),
  }  
}

type TestServer struct{
  *httptest.Server
}

func NewTestServer(t *testing.T, h http.Handler)*TestServer  {
  ts := httptest.NewTLSServer(h)
  return &TestServer{ts}
}

func (ts *TestServer)Get(t *testing.T, urlPath string)(int, http.Header, []byte)  {
  rs, err := ts.Client().Get(ts.URL+urlPath)
  if err != nil{
    t.Fatal(err)
  }
  defer rs.Body.Close() 

  body, err := ioutil.ReadAll(rs.Body)
  if err != nil{
    t.Fatal(err)
  }
  return rs.StatusCode, rs.Header, body
}
