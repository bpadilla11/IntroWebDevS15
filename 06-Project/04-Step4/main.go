/* Brian Padilla
 * 3/14/16
 * PROJECT STEP 4 - refactoring our application, create a new data type called "user" which
 * has fields for the user's name and age. When you receive the user's name and age form submission,
 * create a variable of type "user" then put those values from the form submission into the fields
 * for that variable. Marshal your variable of type "user"  to JSON. Encode that JSON to base64.
 * Store that value in the cookie.
 */

package main

import (
	"fmt"
	"log"
	"html/template"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

func projectWebpage(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	userName := req.FormValue("name")
	userAge := req.FormValue("age")

	currUser := User {
		Name: nameName,
		Age: ageName,
	}

	encode, err := json.Marshal(currentUser)
		if err != nil {
			fmt.Println(err)
		}

	json := base64.StdEncoding.EncodeToString(encode)

	cookie, err := req.Cookie("session")
	id, _ := uuid.NewV4()
	logError(err)

	cookie := &http.Cookie {
		Name:	"session",
		Value: id.String() + userName + userAge,
    // Secure: true,
		HttpOnly:	true,
	}
	http.SetCookie(res, cookie)

	tpl.Execute(res, nil)
}

func main() {
	// set the path URL
	http.HandleFunc("/", projectWebpage)

	fmt.Println("server is now running...") // display when server is running
	http.ListenAndServe(":8080", nil) // set listener to port 8080 on localhost
}
