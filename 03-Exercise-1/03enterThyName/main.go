// Brian Padilla
// 1/31/16
// Simple user name input program

package main

import "fmt"

func main() {
	var userName string // Initalize as a string
	fmt.Print("Enter thy name wary traveler: ")
	fmt.Scan(&userName) // Take user name input
	fmt.Println("Hello", userName) // Output user name
}
