package main

import (
	"log"
	"net/http"

)

func main()  {
  log.Println("Starting main server on port 4000...")
  mux := http.NewServeMux()
  mux.HandleFunc("/", Home)
  mux.HandleFunc("/snippet",ShowSnippet)
  mux.HandleFunc("/snippet/create",CreateSnippet)
  

  if err := http.ListenAndServe(":4000",mux); err != nil{
    log.Fatal(err)
  }
}
