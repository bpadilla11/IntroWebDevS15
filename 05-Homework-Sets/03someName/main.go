/* Brian Padilla
 * 3/5/16
 * Create a webpage that serves at localhost:8080 and will display the name in
 * the url when the url is localhost:8080?n="some-name" - use req.FormValue to do this
 */

package main

import (
	"fmt"
	"io"
	"net/http"
)

// function that handles the website's back end
func thisLittleWebpage(res http.ResponseWriter, req *http.Request) {
	val := req.FormValue("n") // set the FormValue name
	// for simplicity, the webpage is generated here
	page := `
		<!DOCTYPE html>
		<html>
			<head>
				<meta charset="utf-8">
				<title></title>
			</head>
			<body>
				<form method="GET">
					<input type="text" name="n">
					<input type="submit">
				</form>
			</body>
		</html>`
	io.WriteString(res, page + val) // write the page
}

func main() {
	// run the webpage
	http.HandleFunc("/", thisLittleWebpage)

	fmt.Println("server is now running...") // display when server is running
	http.ListenAndServe(":8080", nil) // set listener to port 8080 on localhost
}
