/* Brian Padilla
 * 3/14/16
 * PROJECT STEP 3 - continuing to build our application, create a template which is a form.
 * The form should gather the user's name and age. Store the user's name and age in the cookie.
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
	if err != nil{
		log.Fatalln(err)
	}

	userName := req.FormValue("name")
	userAge := req.FormValue("age")

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
