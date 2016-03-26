/* Brian Padilla
 * 3/25/16
 * Create a web page which serves at localhost over https using TLS
 * NOTE: If the web page states the site is "untrusted", "unsecure", or anything similar,just proceed forward.
 * For more information see 'cert/README.txt'
 */

package main

import (
  "fmt"
  "html/template"
  "log"
  "net/http"
)

// function that handles the website's back end
func thisLittleWebpage(res http.ResponseWriter, req *http.Request) {
  // establish the webpage by parsing the index file
  tpl, err := template.ParseFiles("index.html")
  if err != nil { // if index.html does not exist, give user a error
		log.Fatalln(err) // stops program if file does not exist
	}
	tpl.Execute(res, nil) // execute the html file
}

func main() {
  http.HandleFunc("/", thisLittleWebpage) // set the path URL

  fmt.Println("server is now running...") // display when server is running
  go http.ListenAndServe(":8080", http.RedirectHandler("https://localhost:2525/", 301)) // redirect from HTTP to HTTPS
  http.ListenAndServeTLS(":2525", "cert/cert.pem", "cert/key.pem", nil) // server is now running on HTTPS and load certificates
}
