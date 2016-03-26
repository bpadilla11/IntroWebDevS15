/* Brian Padilla
 * 3/25/16
 * PROJECT STEP 1 - create a web application that serves an HTML template
 */

package main

import (
  "fmt"
	"html/template"
	"log"
	"net/http"
)

// webpage function
// holds all the code for running the website
func projectWebpage(res http.ResponseWriter, req *http.Request) {
  // establish the webpage by parsing the index file
  tpl, err := template.ParseFiles("index.html")
  if err != nil { // if index.html does not exist, give user a error
		log.Fatalln(err) // stops program if file does not exist
	}
	tpl.Execute(res, nil) // execute the html file
}

func main() {
  // set the path URL
	http.HandleFunc("/", projectWebpage)

	fmt.Println("server is now running...") // display when server is running
	http.ListenAndServe(":8080", nil) // set listener to port 8080 on localhost
}
