/* Brian Padilla
 * 3/25/16
 * PROJECT STEP 2 - have the application write a cookie called "session-fino" with a UUID.
 * The cookie should serve HttpOnly and you should have the "Secure" flag set also though
 * comment the "Secure" flag out as we're not using https.
 */

package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"html/template"
	"log"
	"net/http"
)

func projectWebpage(res http.ResponseWriter, req *http.Request) {
	// check if a cookie exists in the current user session
	cookie, err := req.Cookie("session-fino")

	if err != nil { // if no cookie exists, then create one
		id, _ := uuid.NewV4() // generate new id
		cookie = &http.Cookie { // create new session
			Name:	"session-fino", // set session name
			Value: id.String(), // set session id
			// Secure: true, // for HTTPS use, we're not using this
			HttpOnly:	true, // standard HTTP website
		}
		http.SetCookie(res, cookie) // set the generated cookie
		fmt.Println("Created cookie!") // debug message for created cookie
	}

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
