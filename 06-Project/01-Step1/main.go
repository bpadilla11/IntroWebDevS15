/* Brian Padilla
 * 3/14/16
 * PROJECT STEP 1 - create a web application that serves an HTML template
 */

package main

import (
  "fmt"
  "net/http"
  "os"
  "log"
  "text/template"
)

func photoBlog(res http.ResponseWriter, req *http.Request)  {

	templ, err := template.ParseFiles("index.html") // Parse template file
	if err != nil {
		log.Fatalln(err)
	}

	err = templ.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
  http.HandleFunc("/", photoBlog)

  fmt.Println("server is now running...") // display when server is running
  http.ListenAndServe(":8080", nil) // set listener to port 8080 on localhost
}
