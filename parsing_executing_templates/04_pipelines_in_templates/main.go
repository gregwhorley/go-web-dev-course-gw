package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type hero struct {
	Name, Game string
}

type villain struct {
	Name, Game string
}

var tpl *template.Template

var funkMap = template.FuncMap{
	"uc": strings.ToUpper,
	"ff": firstFive,
}

func init() {
	tpl = template.Must(template.New("test1.gohtml").Funcs(funkMap).ParseFiles("templates/test1.gohtml"))
}

func firstFive(s string) string {
	s = strings.TrimSpace(s)
	if len(s) > 5 {
		return s[:5]
	}
	return s
}

func main() {
	a := hero{
		Name: "Mario",
		Game: "Super Mario World",
	}

	b := hero{
		Name: "Link",
		Game: "Legend of Zelda",
	}

	d := villain{
		Name: "Bowser",
		Game: "Super Mario World",
	}

	e := villain{
		Name: "Ganon",
		Game: "Legend of Zelda",
	}

	heroes := []hero{a, b}
	villains := []villain{d, e}

	data := struct {
		GoodGuys []hero
		BadGuys  []villain
	}{
		heroes,
		villains,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "test1.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}
}
