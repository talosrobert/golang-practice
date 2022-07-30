package main

import (
	"flag"
	"net/http"
)

func main() {
	app := newDefaultApplication()

	var pathToCfg string
	flag.StringVar(&pathToCfg, "cfg", "./conf.json", "Configuration file path.")
	flag.Parse()

	cfg := newConfig()
	cfg.loadConfigFromPath(pathToCfg)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fsrv := http.FileServer(http.Dir(cfg.StaticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fsrv))

	app.infoLogger.Printf("Starting server on %s", cfg.getAddressAndPort())

	hsrv := &http.Server{
		Addr:     cfg.getAddressAndPort(),
		Handler:  mux,
		ErrorLog: app.errLogger,
	}

	err := hsrv.ListenAndServe()
	app.errLogger.Fatal(err)
}
