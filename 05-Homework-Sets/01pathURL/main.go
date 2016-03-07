/* Brian Padilla
 * 3/5/16
 * Create a webpage that displays the URL path using req.URL.Path
 */

package main

import (
	"fmt"
	"io"
	"net/http"
)

// prints the URL path to a browser
func thisLittleWebpage(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "<h1>URL: " + req.URL.Path + "</h1>")
}

func main() {
	// set the path URL
	http.HandleFunc("/", thisLittleWebpage)

	fmt.Println("server is now running...") // display when server is running
	http.ListenAndServe(":8080", nil) // set listener to port 8080 on localhost
}
