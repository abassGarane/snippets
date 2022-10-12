package main

import (
	"log"
	"net/http"

	"github.com/abassGarane/snippet/handlers"
)

func main()  {
  log.Println("Starting main server on port 4000...")
  mux := http.NewServeMux()
  mux.HandleFunc("/", handlers.Home)
  mux.HandleFunc("/snippet",handlers.ShowSnippet)
  mux.HandleFunc("/snippet/create",handlers.CreateSnippet)
  

  if err := http.ListenAndServe(":4000",mux); err != nil{
    log.Fatal(err)
  }
}
