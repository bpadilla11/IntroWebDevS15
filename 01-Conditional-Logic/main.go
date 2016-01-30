package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

type paradiseCity struct {
	person
	TakeMeHome bool
}

func main() {
	p1 := paradiseCity {
		person: person {
			Name: "Ben Dover",
			Age: 69,
		},
		TakeMeHome: true,
	}

	tpl, err := template.ParseFiles("templ.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}

}
