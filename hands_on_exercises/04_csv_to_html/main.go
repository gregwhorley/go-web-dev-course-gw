package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)

/*
1. Parse table.csv CSV file, putting two fields from the contents of the CSV file into a data structure.
2. Parse an html template, then pass the data from step 1 into the CSV template; have the html template display the CSV data in a web page.
*/

type Record struct {
	Date  time.Time
	Stock float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	fh, err := os.Open("table.csv")
	defer fh.Close()
	if err != nil {
		log.Fatalln(err)
	}
	r := csv.NewReader(fh)
	lines, _ := r.ReadAll()
	records := make([]Record, 0, len(lines))
	for i, line := range lines {
		if i == 0 {
			continue
		}
		date, _ := time.Parse("2006-01-02", line[0])
		open, _ := strconv.ParseFloat(line[1], 64)
		records = append(records, Record{
			Date: date, Stock: open,
		})
	}
	tplErr := tpl.Execute(os.Stdout, records)
	if tplErr != nil {
		log.Fatalln(tplErr)
	}
}
