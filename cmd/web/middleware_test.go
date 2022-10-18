package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecureHeaders(t *testing.T)()  {
  // Running tests in parallel
  t.Parallel()
  rr := httptest.NewRecorder()

  r, err := http.NewRequest("GET", "/", nil)
  if err != nil{
    t.Fatal(err)
  }

  next := http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
    w.Write([]byte("ok"))
  })
  secureHeaders(next).ServeHTTP(rr,r)

  // Testing for X-Frame-Options header
  rs := rr.Result() 
   framesOptions := rs.Header.Get("X-Frame-Options")
  if framesOptions != "deny"{
    t.Errorf("want %q; got %q", "deny", framesOptions)
  }

  // tests for x-xss protection header
  xssProtection := rs.Header.Get("X-XSS-Protection")
  if xssProtection != "1;mode=block"{
    t.Errorf("want %q; got %q", "1;mode=block", xssProtection)
  } 

  // status code
  if rs.StatusCode != http.StatusOK {
    t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
  }
  defer rs.Body.Close() 
  body, err := ioutil.ReadAll(rs.Body)
  if err != nil{
    t.Fatal(err)
  }
  if string(body) != "ok"{
    t.Errorf("want the body to equal %q", "ok")
  }
}
