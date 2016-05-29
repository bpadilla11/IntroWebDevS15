/* Brian Padilla
 * 3/21/16
 * PROJECT STEP 8 - Allow the user to logout.
 * Show a log-in button when the user is not logged-in.
 * Show a log-out button only when the user is logged in.
 */

package main

import (
	"crypto/hmac"
	"crypto/sha256"
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

// secure data using an HMAC method
func secureData(data string) string {
	h := hmac.New(sha256.New, []byte(data + "SuperCerealSecurity")) // secure incoming data
	return string(h.Sum(nil)) // return secured data
}

// function that handles the website's back end
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
			Name: "session-fino", // set session name
			Value: encode, // set session id
			// Secure: true, // for HTTPS use, we're not using this
			HttpOnly: true, // standard HTTP website
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
		secure := secureData(userValues.UUID + userValues.Name + userValues.Age) // secure the data from modification
		b, _ := json.Marshal(secure) // convert back to json
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
