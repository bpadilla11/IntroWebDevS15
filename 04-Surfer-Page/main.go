// Brian Padilla
// 2/2/16
// A Go program that hosts a webpage on a local machine

/*
	How to run:
	1. Open the destination of this file in a terminal.
	2. Type "go run main.go" to run the program. Continue running it in terminal background.
	3. Open a web browser.
	4. Type "127.0.0.1:8080" in the address bar.
	5. The surfer webpage should successfully appear in the browser.
	6. Close the broswer once complete.
	7. Go back to the terminal and press "ctrl+c" close the program.
*/

package main

import (
	"html/template"
	"net/http"
	"log"
)

// Function to establish a parser for the index.html file
func surfsUpDude(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil { // If index.html does not exist, give user a error
		log.Fatalln(err)
	}
	// Execute the html file
	tpl.Execute(res, nil)
}

func main() {
	http.HandleFunc("/", surfsUpDude) // Handles the main page
	http.Handle("/css/",http.StripPrefix("/css",http.FileServer(http.Dir("css")))) // Handles the css file
	http.Handle("/pic/",http.StripPrefix("/pic",http.FileServer(http.Dir("pic")))) // Handles the pic file
	http.ListenAndServe(":8080", nil) // ListenAndServe on local ip with port 8080
}
