package main

import "net/http"

func (app *application) routes(staticDirPath string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fsrv := http.FileServer(http.Dir(staticDirPath))
	mux.Handle("/static/", http.StripPrefix("/static", fsrv))

	return mux
}
