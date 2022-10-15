package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/abassGarane/snippet/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
)

type application struct{
  errorLog *log.Logger
  infoLog *log.Logger
  snippets *mysql.SnippetModel
  templateCache map[string]*template.Template
  session *sessions.Session
}

func main()  {
  //Parsing commandline flags
  addr := flag.String("addr",":4000","HTTP network address")
  dsn := flag.String("dsn","root:philos@/snippetbox?parseTime=true","Mysql database Connection string")
  secret := flag.String("secret", "kayfdas8saiXBUusu6ra", "Secret for session management")
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

  // Initialize a template cache
  templateCache, err := newTemplateCache("./ui/html/")
  if err != nil{
    errLog.Fatal(err)
  }

  // Initialize sessions
  session := sessions.New([]byte(*secret))
  session.Lifetime = 12*time.Hour

  app := &application{
    errorLog: errLog,
    infoLog: infoLog,
    session: session,
    snippets: &mysql.SnippetModel{DB: db},
    templateCache: templateCache,
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
  db.SetMaxOpenConns(95)
  db.SetMaxIdleConns(5)
  return db, nil
}
