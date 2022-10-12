package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request)  {
  if r.URL.Path != "/"{
    log.Println("NotFound page called")
    http.NotFound(w,r)
    return
  }
  log.Println("Home snippet handler called...")
  w.Write([]byte("Hello from home"))
}

func ShowSnippet( w http.ResponseWriter, r *http.Request)  {
  log.Println("Show snippet handler called...")
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1{
    http.NotFound(w,r)
    return
  }
  // w.Write([]byte("Display a specific snippet ..."))
  fmt.Fprintf(w,"Display a specific snippet with id %d...",id)
}

func CreateSnippet( w http.ResponseWriter, r *http.Request)  {
  log.Println("CreateSnippet snippet handler called...")
  // Removing a header
  w.Header()["Date"]=nil
  if r.Method != "POST"{
    w.Header().Set("Allow","POST")
    // w.WriteHeader(http.StatusMethodNotAllowed)
    // w.Write([]byte("Method not allowed"))
    http.Error(w,"Method not Allowed", http.StatusMethodNotAllowed)
    return
  }
  // Update or modify an existing header
  w.Header().Set("Content-Type","application/json")
  w.Write([]byte(`{"name":"abass garane"}`))
}

