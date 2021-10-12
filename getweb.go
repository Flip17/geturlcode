package main

// +ignore build
import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func renderTemp(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("form.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("Error executing html :", err)
		return
	}
}

func showurl(w http.ResponseWriter, r *http.Request) {
	var linkurl = r.FormValue("geturl")
	resp, err := http.Get(linkurl)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)

	t, err := template.ParseFiles("form.html")
	links := struct {
		Url string
	}{
		Url: sb,
	}
	t.Execute(w, links)

}

func main() {

	http.HandleFunc("/", renderTemp)
	http.HandleFunc("/geturl", showurl)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
