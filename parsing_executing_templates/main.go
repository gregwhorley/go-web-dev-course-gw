package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Recipient struct {
	Name, Email, Referral string
	Interested            bool
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	recipients := []Recipient{
		{"Aardvark", "aardvark@email.com", "", false},
		{"Billy Bob", "bthornton@hotmail.com", "ajolie@hotmail.com", true},
		{"Chris Rock", "chris.rock@gmail.com", "", true},
	}
	for _, r := range recipients {
		err := tpl.Execute(os.Stdout, r)
		if err != nil {
			log.Fatal(err)
		}
	}
}
