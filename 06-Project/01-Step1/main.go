/* Brian Padilla
 * 3/14/16
 * PROJECT STEP 1 - create a web application that serves an HTML template
 */

package main

import(
  "fmt"
	"log"
	"html/template"
	"net/http"
)

func projectWebpage(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil{
		log.Fatalln(err)
	}
	tpl.Execute(res, nil)
}

func main() {
  // set the path URL
	http.HandleFunc("/", projectWebpage)

	fmt.Println("server is now running...") // display when server is running
	http.ListenAndServe(":8080", nil) // set listener to port 8080 on localhost
}
