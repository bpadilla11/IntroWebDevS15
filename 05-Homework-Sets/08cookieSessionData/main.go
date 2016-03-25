/* Brian Padilla
 * 3/25/16
 * Create a webpage which writes a cookie to the client's machine.
 * Though this is NOT A BEST PRACTICE, you will store some session data in the cookie.
 * Make sure you use HMAC to ensure that session data is not changed by a user.
 */

package main

import (
  "crypto/hmac"
  "crypto/sha256"
  "fmt"
  "github.com/nu7hatch/gouuid"
  "html/template"
  "io"
  "log"
  "net/http"
  "strings"
)

// function that encrypts input data using a SHA256 encryption
func getData(data string) string {
	h := hmac.New(sha256.New, []byte("be#r=*resw6rEjez")) // create new SHA256 key
	io.WriteString(h, data) // write key as a string
	return fmt.Sprintf("%x", h.Sum(nil)) // return the encrypted key
}

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

  if req.Method == "POST" { // if a cookie exists and user submits data, then update cookie
		newValue := req.FormValue("cookieType") // get user name and age from the index.html file
    userCookie := getData(newValue) // encrypt the user input
		tempValue := strings.Split(cookie.Value," | ") // split the session id and newValue
		cookie.Value = tempValue[0] + " | " + userCookie // keep current session id from tempValue, but update user name and age
		tempValue = nil // not necessary, but thought the tempValue array should be cleared of the id
		http.SetCookie(res, cookie) // set the updated cookie
		fmt.Println("Updated cookie!") // debug message for updated cookie
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
