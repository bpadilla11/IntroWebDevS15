/* Brian Padilla
 * 3/5/16
 * Create a webpage that serves at localhost:8080 and will display the name in
 * the url when the url is localhost:8080/name - use req.URL.Path to do this
 */

package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// function that handles the website's back end
func thisLittleWebpage(res http.ResponseWriter, req *http.Request) {
	name := strings.Split(req.URL.Path, "/")
	nameOutput := strings.Join(name, "/")
	io.WriteString(res, "<h1>URL: " + nameOutput + "</h1>")
}

func main() {
	/// set the path URL
	http.HandleFunc("/", thisLittleWebpage)

	fmt.Println("server is now running...") // display when server is running
	http.ListenAndServe(":8080", nil) // set listener to port 8080 on localhost
}
