package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{
		l: l,
	}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello, friend.")

	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "faaaaaaaailed", http.StatusBadRequest)
	}

	fmt.Fprintf(rw, "Hello %s from %s\n", d, r.RemoteAddr)
}
