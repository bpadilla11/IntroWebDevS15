/* Brian Padilla
 * 3/25/16
 * PROJECT STEP 4 - refactoring our application, create a new data type called "user" which
 * has fields for the user's name and age. When you receive the user's name and age form submission,
 * create a variable of type "user" then put those values from the form submission into the fields
 * for that variable. Marshal your variable of type "user" to JSON. Encode that JSON to base64.
 * Store that value in the cookie.
 */

package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	UUID, Name, Age string
}

func projectWebpage(res http.ResponseWriter, req *http.Request) {
	// check if a cookie exists in the current user session
	cookie, err := req.Cookie("session-fino")

	if err != nil { // if no cookie exists, then create one
		id, _ := uuid.NewV4() // generate new id
		userSess := User {
			UUID: id.String(), // put 'uuid' in json
		}
		b, _ := json.Marshal(userSess) // marshal 'userSess' uuid into json.
		encode := base64.StdEncoding.EncodeToString(b) // encode json to base64
		cookie = &http.Cookie { // create new session
			Name:   "session-fino", // set session name
			Value: encode, // set session id
			// Secure: true, // for HTTPS use, we're not using this
			HttpOnly:   true, // standard HTTP website
		}
		http.SetCookie(res, cookie) // set the generated cookie
		fmt.Println("Created cookie!") // debug message for created cookie
	}

	if req.Method == "POST" { // if a cookie has been generated and updated, then update the cookie
		decode, _ := base64.StdEncoding.DecodeString(cookie.Value) // recover the json
		var userValues User // for keeping the unmarshalled json
		json.Unmarshal(decode, &userValues) // unmarshal the json
		userValues.Name = req.FormValue("userName") // get userName input
		userValues.Age = req.FormValue("userAge") // get userAge input
		b, _ := json.Marshal(userValues) // convert back to json
		encode := base64.StdEncoding.EncodeToString(b) // encode the json to base64
		cookie.Value = encode // update new values in the cookie
		http.SetCookie(res, cookie) // set the updated cookie with same uuid, but different name and age
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
