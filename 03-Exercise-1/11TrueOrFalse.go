// Brian Padilla
// 1/31/16
// Finds the given boolean example

package main

import "fmt"

func main() {
	whatIsValue := (true && false) || (false && true) || !(false && false)
	fmt.Println(whatIsValue) // Outputs 'true'
}
