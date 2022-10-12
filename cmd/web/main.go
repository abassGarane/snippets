package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main()  {
  //Parsing commandline flags
  addr := flag.String("addr",":4000","HTTP network address")
  flag.Parse()
  
  // custom logger
  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
  errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
  infoLog.Printf("Starting main server on port %s...\n",*addr)
  fileServer := http.FileServer(http.Dir("./ui/static"))

  mux := http.NewServeMux()
  mux.HandleFunc("/", Home)
  mux.HandleFunc("/snippet",ShowSnippet)
  mux.HandleFunc("/snippet/create",CreateSnippet)
  mux.Handle("/static/", http.StripPrefix("/static",fileServer))
  
  srv := &http.Server{
    Addr: *addr,
    ErrorLog: errLog,
    Handler: mux,
  }

  if err := srv.ListenAndServe(); err != nil{
    errLog.Fatal(err)
  }
}
