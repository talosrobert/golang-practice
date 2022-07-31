package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/talosrobert/golang-practice/pkg/models/mysql"
)

type application struct {
	infoLogger *log.Logger
	errLogger  *log.Logger
	snippets   *mysql.SnippetModel
}

func newDefaultApplication() *application {
	return &application{
		infoLogger: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errLogger:  log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		snippets:   &mysql.SnippetModel{},
	}
}

func (app *application) setDatabase(db *sql.DB) {
	app.snippets.DB = db
}
