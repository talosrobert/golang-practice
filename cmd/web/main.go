package main

import (
	"log"
	"net/http"
	"flag"
	"os"
)

type application struct {
	infoLogger *log.Logger
	errLogger *log.Logger
}

type config struct {
	Addr string
	StaticDir string
}

func newConfig() *config {
	return &config{}
}

func main() {
	app := &application{
		infoLogger: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errLogger: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	cfg := newConfig()
	flag.StringVar(&cfg.Addr, "addr", ":8080", "HTTP Network address")
	flag.StringVar(&cfg.StaticDir, "staticdir", "./ui/static", "Static files directory")
	flag.Parse()

	router := http.NewServeMux()
	router.HandleFunc("/", app.home)
	router.HandleFunc("/snippet", app.showSnippet)
	router.HandleFunc("/snippet/create", app.createSnippet)

	fsvr := http.FileServer(http.Dir(cfg.StaticDir))
	router.Handle("/static/", http.StripPrefix("/static", fsvr))

	app.infoLogger.Printf("Starting server on %s", cfg.Addr)

	hsvr := &http.Server{
		Addr: cfg.Addr,
		Handler: router,
		ErrorLog: app.errLogger,
	}

	err := hsvr.ListenAndServe()
	app.errLogger.Fatal(err)
}
