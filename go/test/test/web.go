package main

import (
	//	"log"
	"net/http"
	//	"html/template"
	//	"path/filepath"
	"bufio"
	"fmt"
	"os"
)

func fileOpen(w http.ResponseWriter, r *http.Request) {
	fp, err := os.Open("index/hoge.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	sccaner := bufio.NewScanner(fp)
	for sccaner.Scan() {
		fmt.Println(sccaner.Text())
	}
}

func main() {

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("gol"))
	mux.Handle("/gol/", http.StripPrefix("/gol/", files))

	//	mux.HandleFunc("/",fileOpen)

	server := &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: mux,
	}
	server.ListenAndServe()
}

/*
func index(w http.ResponseWriter, r *http.Request){
	files:=[]string{
		"template/layout.html",
		"template/nevbar.html",
		"template/index.html",
	}
	templates:=template.Must(template.ParseFiles(files...))
	threads,
}
*/
