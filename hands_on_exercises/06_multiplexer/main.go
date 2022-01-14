package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cheese", cheese)
	http.HandleFunc("/", hello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func cheese(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Cheese is good</strong></body></html>`)
}

func hello(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`)
}
