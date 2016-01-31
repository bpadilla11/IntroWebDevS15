/*
	Brian Padilla
	1/30/15
	Go template use with a reference to Paradise City by Guns 'n Roses.
*/

package main

import (
	"log"
	"os"
	"text/template"
)

// Person struct: takes a Name as a string and Age as an int
type person struct {
	Name string
	Age  int
}

// Paradise City struct: takes from the person struct and TakeMeHome as a bool
type paradiseCity struct {
	person
	TakeMeHome bool
}

func main() {
	p1 := paradiseCity { // Create a Paradise City
		person: person { // Create a person
			Name: "Slash", // Best guitar player of all time!
			Age: 50, // Age of Slash as of 1/30/16
		},
		TakeMeHome: true, // Oh won't you please take me home!
	}

	// Parse the templ.gohtml file
	tpl, err := template.ParseFiles("templ.gohtml")
	if err != nil { // If templ.gohtml does not exist, give user a error
		log.Fatalln(err)
	}

	// Execute the webpage
	err = tpl.Execute(os.Stdout, p1)
	if err != nil { // If p1 does not exist, give user a error
		log.Fatalln(err)
	}

}
