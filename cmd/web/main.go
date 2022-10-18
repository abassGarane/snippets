package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/abassGarane/snippet/pkg/models"
	"github.com/abassGarane/snippet/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
)

type application struct{
  errorLog *log.Logger
  infoLog *log.Logger

  snippets interface{
    Insert(string,string,string)(int,error)
    Get(int)(*models.Snippet,error)
    Latest()([]*models.Snippet, error)
  }
  templateCache map[string]*template.Template
  session *sessions.Session

  users interface{
    Insert(string,string,string)error
    Authenticate(string, string)(int,error)
    Get(int)(*models.User, error)
  }
}
//context keys
type contextKey string
var contextKeyUser = contextKey("user")

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
  session.SameSite=http.SameSiteStrictMode
  session.Secure = true

  app := &application{
    errorLog: errLog,
    infoLog: infoLog,
    session: session,
    snippets: &mysql.SnippetModel{DB: db},
    templateCache: templateCache,
    users: &mysql.UserModel{DB: db},
  }

  // tls config
  tlsConfig := &tls.Config{
    PreferServerCipherSuites: true,
    CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
  }

  infoLog.Printf("Starting main server on port %s...\n",*addr)
    
  srv := &http.Server{
    Addr: *addr,
    ErrorLog: errLog,
    Handler: app.Routes(),
    TLSConfig: tlsConfig,
    IdleTimeout: time.Minute,
    ReadTimeout: 5 *time.Second,
    WriteTimeout: 10 *time.Second,
  }

  if err := srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem"); err != nil{
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
