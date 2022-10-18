package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
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
  jar, err:= cookiejar.New(nil)
  if err != nil{
    t.Fatal(err)
  }

  //automatically store cookies
  ts.Client().Jar = jar

  //Disable redirects
  ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error { 
    return http.ErrUseLastResponse
  }

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
