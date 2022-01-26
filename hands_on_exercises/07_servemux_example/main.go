package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type HttpData struct {
	Method        string
	URL           *url.URL
	Submissions   map[string][]string
	Header        http.Header
	Host          string
	ContentLength int64
}

func main() {
	http.Handle("/", http.HandlerFunc(entryPoint))
	http.Handle("/dog", http.HandlerFunc(dogRoute))
	http.Handle("/me", http.HandlerFunc(meRoute))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func meRoute(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello Me!")
}

func dogRoute(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello Dog!")
}

func entryPoint(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := HttpData{
		Method:        request.Method,
		URL:           request.URL,
		Submissions:   request.Form,
		Header:        request.Header,
		Host:          request.Host,
		ContentLength: request.ContentLength,
	}

	tpl.ExecuteTemplate(writer, "index.gohtml", data)
}
