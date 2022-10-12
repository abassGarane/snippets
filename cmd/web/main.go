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
    
  srv := &http.Server{
    Addr: *addr,
    ErrorLog: errLog,
    Handler: app.Routes(),
  }

  if err := srv.ListenAndServe(); err != nil{
    errLog.Fatal(err)
  }
}
