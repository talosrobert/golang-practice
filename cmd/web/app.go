package main

import (
	"log"
	"os"
)

type application struct {
	infoLogger *log.Logger
	errLogger  *log.Logger
}

func newDefaultApplication() *application {
	return &application{
		infoLogger: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errLogger:  log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
