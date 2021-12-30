package main

import (
	"log"
	"os"
	"text/template"
)

/*
1. Create a data structure to pass to a template which
* contains information about California hotels including Name, Address, City, Zip, Region
* region can be: Southern, Central, Northern
* can hold an unlimited number of hotels
*/

type Hotel struct {
	Region, Name, Address, City, Zip string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []Hotel{
		Hotel{
			Region:  "Southern",
			Name:    "Eagles Hotel",
			Address: "123 Lovely Pl",
			City:    "Los Angeles",
			Zip:     "98765",
		},
		Hotel{
			Region:  "Central",
			Name:    "Hilton Central",
			Address: "987 Paris Blvd",
			City:    "Mendocino",
			Zip:     "47474",
		},
		Hotel{
			Region:  "Northern",
			Name:    "Four Seasons",
			Address: "444 Dude St",
			City:    "San Bernadino",
			Zip:     "22333",
		},
	}
	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
