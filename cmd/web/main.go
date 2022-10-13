package main

import (
  "database/sql"
	"flag"
	"log"
	"net/http"
	"os"
  _ "github.com/go-sql-driver/mysql"
)

type application struct{
  errorLog *log.Logger
  infoLog *log.Logger
}

func main()  {
  //Parsing commandline flags
  addr := flag.String("addr",":4000","HTTP network address")
  dsn := flag.String("dsn","root:philos@/snippetbox?parseTime=true","Mysql database Connection string")
  flag.Parse()
  
  // custom logger
  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
  errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

  //intialization of database
  db, err := openDB(*dsn)
  if err != nil{
    errLog.Fatal(err)
  }
  defer db.Close()

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

func openDB(dsn string)(*sql.DB, error)  {
  db, err := sql.Open("mysql", dsn)
  if err != nil{
    return nil, err
  }
  if err = db.Ping(); err != nil{
    return nil, err
  }
  return db, nil
}
