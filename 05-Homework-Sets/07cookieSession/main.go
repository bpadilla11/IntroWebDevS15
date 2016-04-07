/* Brian Padilla
 * 3/25/16
 * Create a webpage which writes a cookie to the client's machine.
 * This cookie should be designed to create a session and should use a UUID,
 * HttpOnly, and Secure (though you'll need to comment secure out).
 */

package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"html/template"
	"log"
	"net/http"
)

// function that handles the website's back end
func thisLittleWebpage(res http.ResponseWriter, req *http.Request) {
	// load cookie if it exists
	cookie, err := req.Cookie("session")
	if err != nil { // if no cookie exists, then create one
		id, _ := uuid.NewV4() // generate new id
		cookie = &http.Cookie { // create new session
			Name:	"session", // set session name
			Value: id.String(), // set session id
			//Secure: true, // for HTTPS use, we're not using this
			HttpOnly:	true, // standard HTTP website
		}
		http.SetCookie(res, cookie) // set the generated cookie
		fmt.Println("Created cookie!") // debug message for created cookie
	}

  tpl, err := template.ParseFiles("index.html")
  if err != nil { // if index.html does not exist, give user a error
    log.Fatalln(err) // stops program if file does not exist
  }

  tpl.Execute(res, nil) // execute the html file
}

func main() {
  // set the path URL
  http.HandleFunc("/", thisLittleWebpage)

  fmt.Println("server is now running...") // display when server is running
  http.ListenAndServe(":8080", nil) // set listener to port 8080 on localhost
}
