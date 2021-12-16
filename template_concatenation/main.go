package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Stu Dent"
	title := "Golang person"
	description := "The course is called templating with concatenation but this is also string interpolation..."
	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	<h2>` + title + `</h2>
	<p>`+ description +`</p>
	</body>
	</html>
	`
	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer newFile.Close()
	io.Copy(newFile, strings.NewReader(tpl))
}
