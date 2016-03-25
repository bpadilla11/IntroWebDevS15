/* Brian Padilla
 * 3/6/16
 * Create a webpage which uses a cookie to track the number of visits of a user.
 * Display the number of visits. Make sure that the favicon.ico requests are not
 * also incrementing the number of visits
 */

 package main

import (
  "fmt"
 	"io"
 	"net/http"
 	"strconv"
)

// function that handles the website's back end
func thisLittleWebpage(res http.ResponseWriter, req *http.Request) {
  // no favicon requests
  if req.URL.Path != "/" {
    http.NotFound(res, req)
    return
  }
  // try to get the cookie
  cookie, err := req.Cookie("my-cookie")
  // if there is no cookie, then create one
  if err == http.ErrNoCookie {
    cookie = &http.Cookie{
      Name:  "my-cookie",
      Value: "0",
    }
  }
  // there is always a cookie at this point
  count, _ := strconv.Atoi(cookie.Value)
  count++ // increment page visits
  cookie.Value = strconv.Itoa(count) // convert value to string

  http.SetCookie(res, cookie)

  io.WriteString(res, cookie.Value)
}

func main() {
  http.HandleFunc("/", thisLittleWebpage)

  fmt.Println("server is now running...") // display when server is running
  http.ListenAndServe(":8080", nil) // set listener to port 8080 on localhost
}
