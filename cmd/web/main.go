package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct{
  errorLog *log.Logger
  infoLog *log.Logger
}

func main()  {
  //Parsing commandline flags
  addr := flag.String("addr",":4000","HTTP network address")
  flag.Parse()
  
  // custom logger
  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
  errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

  app := &application{
    errorLog: errLog,
    infoLog: infoLog,
  }

  infoLog.Printf("Starting main server on port %s...\n",*addr)
  fileServer := http.FileServer(http.Dir("./ui/static"))

  mux := http.NewServeMux()
  mux.HandleFunc("/", app.Home)
  mux.HandleFunc("/snippet",app.ShowSnippet)
  mux.HandleFunc("/snippet/create",app.CreateSnippet)
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
