package main

import (
	//"text/template"
	"html/template"
	"os"
)

type Person struct {
	Name    string
	Age     int
	Hobbies []string
}

var funcs = template.FuncMap{
	"increment": func(num int) int {
		return num + 1
	},
}

func main() {

	ray := &Person{"Ray", 35, []string{"Read", "Paint", "Math and Vocabulary puzzles"}}
	// LoadTemplate("greetings.txt", ray)
	// LoadTemplate("greetings2.txt", ray)
	LoadTemplate("index.html", ray)
}

func LoadTemplate(fileName string, data interface{}) {
	t, err := template.New(fileName).Funcs(funcs).ParseFiles("templates/" + fileName)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
