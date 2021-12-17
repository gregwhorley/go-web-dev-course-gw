package main

import (
	"log"
	"os"
	"text/template"
)

type hero struct {
	Name, Game string
}

type villain struct {
	Name, Game string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
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

	c := hero{
		Name: "Samus",
		Game: "Metroid",
	}

	d := villain{
		Name: "Bowser",
		Game: "Super Mario World",
	}

	e := villain{
		Name: "Ganon",
		Game: "Legend of Zelda",
	}

	f := villain{
		Name: "Mother Brain",
		Game: "Metroid",
	}

	heroes := []hero{a, b, c}
	villains := []villain{d, e, f}

	data := struct {
		GoodGuys []hero
		BadGuys  []villain
	}{
		heroes,
		villains,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
