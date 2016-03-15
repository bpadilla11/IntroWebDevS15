/* Brian Padilla
 * 3/14/16
 * PROJECT STEP 2 - have the application write a cookie called "session-fino" with a UUID.
 * The cookie should serve HttpOnly and you should have the "Secure" flag set also though
 * comment the "Secure" flag out as we're not using https.
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

	cookie, err := req.Cookie("session")

	id, _ := uuid.NewV4()
	logError(err)

	cookie := &http.Cookie {
		Name:	"session",
		Value: id.String(),
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
