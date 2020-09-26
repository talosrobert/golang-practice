package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/talosrobert/golang-practice/cmd/api/handlers"
)

func main() {
	l := log.New(os.Stdout, "hello-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	bh := handlers.NewBye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/bye", bh)

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
