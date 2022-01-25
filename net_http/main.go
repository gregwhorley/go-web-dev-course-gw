package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type gregTest int

type HttpData struct {
	Method        string
	URL           *url.URL
	Submissions   map[string][]string
	Header        http.Header
	Host          string
	ContentLength int64
}

func (m gregTest) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Greg-Type", "this is a custom header...")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := HttpData{
		Method:        req.Method,
		URL:           req.URL,
		Submissions:   req.Form,
		Header:        req.Header,
		Host:          req.Host,
		ContentLength: req.ContentLength,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d gregTest
	log.Fatal(http.ListenAndServe(":8080", d))
}
