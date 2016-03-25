/* Brian Padilla
 * 3/5/16
 * Create a webpage that serves a form and allows the user to enter their name.
 * Once a user has entered their name, show their name on the webpage.
 * Use req.FormValue to do this
 */

package main

import (
	"fmt"
	"io"
	"net/http"
)

// function that handles the website's back end
func thisLittleWebpage(res http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name") // set the FormValue name
	// for simplicity, the webpage is generated here
	page := `
		<!DOCTYPE html>
		<html>
		  <head>
		    <meta charset="utf-8">
		    <title></title>
		  </head>
		  <body>
		    <form method="GET"> Input your name:
		      <input type="text" name="name">
		      <input type="submit">
		    </form>
		  </body>
		</html>`

	io.WriteString(res, page + name) // write the page
}

func main() {
	// run the webpage
	http.HandleFunc("/", thisLittleWebpage)

	fmt.Println("server is now running...") // display when server is running
	http.ListenAndServe(":8080", nil) // set listener to port 8080 on localhost
}
