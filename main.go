package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	/*		"os"
	"bufio"

	"html/template"
	*/)

func wordOp(w http.ResponseWriter, r *http.Request) {
	e := r.ParseForm()
	log.Println(e)
	res := r.FormValue("name")
	log.Println(res)
	fmt.Fprintf(w, res)
	/*	temp_file := []string{"statics/index.html"}
		var templates *template.Template
		templates = template.Must(template.)
	*/
}

func getNews(w http.ResponseWriter, r *http.Request) {
	type Artl []struct {
		Title string `json:"title"`
	}

	type Obj struct {
		//	TotalResults string `json:"totalResults"`
		Author string `json:"author"`
		//title	string 'json:"articles:title"'
		Status   string `json:"status"`
		Articles Artl   `json:"articles"`
	}

	que := r.FormValue("que")
	var reqURL string
	var key string = "&apikey=6b07c113972b48b49a032ce525b9a7e3"

	reqURL += "https://newsapi.org/v2/top-headlines?" + "q=" + que + "&country=us" + key

	res, err := http.Get(reqURL)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[status]%d¥n", res.StatusCode)

	defer res.Body.Close()
	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		log.Fatal(error)
	}
	//fmt.Printf(string(body))

	var objs Obj
	/*
		err := json.Unmarshal(body, &objs)
		if err != nil{
			Sorry:="sorry"
			t.Execute(w,Sorry)
		}
	*/
	error = json.Unmarshal(body, &objs)
	if error != nil {
		log.Fatal(error)
	}

	log.Printf("%+v", objs.Status)
	//	log.Printf("%+v",objs.Articles[0].Title)

	tmpl_files := []string{
		"statics/tmpl.html",
		"statics/index.html"}
	var templates *template.Template

	//t,_ := template.ParseFiles("statics/tmpl.html")
	//templates = template.Must(template.ParseFiles(tmpl_files...))

	if len(objs.Articles) != 0 {

		for _, p := range objs.Articles {
			//	t.Execute(w,string(p.Title))
			//templates.Execute(w, string(p.Title))
			templates = template.Must(template.ParseFiles(tmpl_files...))
			templates.ExecuteTemplate(w, "tmpl", string(p.Title))
		}
	} else {
		templates.ExecuteTemplate(w, "tmpl", string("Not Any Result"))
	}

}

func main() {

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("statics"))

	mux.Handle("/app", http.StripPrefix("/app", files))
	mux.HandleFunc("/lang", wordOp)
	mux.HandleFunc("/news", getNews)

	//	mux.HandleFunc("/",fileOpen)

	server := &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: mux,
	}
	server.ListenAndServe()
}
