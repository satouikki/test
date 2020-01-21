package main

import (
	"net/http"
	pkg "pkg"
)

func main() {

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("statics"))

	mux.Handle("/app", http.StripPrefix("/app", files))
	mux.HandleFunc("/news", pkg.GetNews)

	server := &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: mux,
	}
	server.ListenAndServe()
}
