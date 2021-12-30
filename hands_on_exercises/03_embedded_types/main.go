package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip string
}

type region struct {
	Hotel  hotel
	Region string
}

type Regions []region

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r := Regions{
		region{
			Hotel: hotel{
				Name:    "Hotel Transylvania",
				Address: "6-5000 Spooky St",
				City:    "San Francisco",
				Zip:     "29384",
			},
			Region: "Southern",
		},
		region{
			Hotel: hotel{
				Name:    "Hotel Cali Bro",
				Address: "4488 Dude Ln",
				City:    "Berkeley",
				Zip:     "44885",
			},
			Region: "Central",
		},
		region{
			Hotel: hotel{
				Name:    "Hilton",
				Address: "48484 Hamburger St",
				City:    "San Bernadino",
				Zip:     "88999",
			},
			Region: "Northern",
		},
	}

	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}
