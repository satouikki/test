package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

func GetNews(w http.ResponseWriter, r *http.Request) {

	type Artl []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
	}

	type Obj struct {
		TotalResults int    `json:"totalResults"`
		Author       string `json:"author"`
		Status       string `json:"status"`
		Articles     Artl   `json:"articles"`
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

	var objs Obj
	error = json.Unmarshal(body, &objs)
	if error != nil {
		log.Fatal(error)
	}

	log.Printf("%+v", objs.Status)

	tmpl_files := []string{
		"statics/tmpl.html",
		"statics/index.html",
	}

	var templates *template.Template

	if len(objs.Articles) > 0 {

		p := objs.Articles

		templates = template.Must(template.ParseFiles(tmpl_files...))
		templates.ExecuteTemplate(w, "tmpl", p)

	} else {
		str := "Not Any Result"

		templates = template.Must(template.ParseFiles(tmpl_files...))
		templates.ExecuteTemplate(w, "tmpl", str)
	}
}
