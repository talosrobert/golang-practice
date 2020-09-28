package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/talosrobert/golang-practice/cmd/api/handlers"
)

func main() {
	l := log.New(os.Stdout, "hello-api", log.LstdFlags)
	ph := handlers.NewProduct(l)

	sm := mux.NewRouter()
	getSubrouter := sm.Methods("GET").Subrouter()
	getSubrouter.HandleFunc("/", ph.Get)

	putSubrouter := sm.Methods("PUT").Subrouter()
	putSubrouter.HandleFunc("/{id:[0-9]+}", ph.Update)
	putSubrouter.Use(ph.MiddlewareProductValidation)

	postSubrouter := sm.Methods("POST").Subrouter()
	postSubrouter.HandleFunc("/", ph.Add)
	postSubrouter.Use(ph.MiddlewareProductValidation)

	s := &http.Server{
		Addr:         ":8083",
		Handler:      sm,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	fmt.Println("starting graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
