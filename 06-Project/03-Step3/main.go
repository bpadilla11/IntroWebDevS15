/* Brian Padilla
 * 3/25/16
 * PROJECT STEP 3 - continuing to build our application, create a template which is a form.
 * The form should gather the user's name and age. Store the user's name and age in the cookie.
 */

package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func projectWebpage(res http.ResponseWriter, req *http.Request) {
	// load cookie if it exists
	cookie, err := req.Cookie("session-fino")
	if err != nil { // if no cookie exists, then create one
		id, _ := uuid.NewV4() // generate new id
		cookie = &http.Cookie { // create new session
			Name:	"session-fino", // set session name
			Value: id.String(), // set session id
			//Secure: true, // for HTTPS use, we're not using this
			HttpOnly:	true, // standard HTTP website
		}
		http.SetCookie(res, cookie) // set the generated cookie
		fmt.Println("Created cookie!") // debug message for created cookie
	}

	if req.Method == "POST" { // if a cookie exists and user submits data, then update cookie
		newValue := req.FormValue("userName") + " " + req.FormValue("userAge") // get user name and age from the index.html file
		tempValue := strings.Split(cookie.Value," | ") // split the session id and newValue
		cookie.Value = tempValue[0] + " | " + newValue // keep current session id from tempValue, but update user name and age
		tempValue = nil // not necessary, but thought the tempValue array should be cleared of the id
		http.SetCookie(res, cookie) // set the updated cookie
		fmt.Println("Updated cookie!") // debug message for updated cookie
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
