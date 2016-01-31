// Brian Padilla
// 1/31/16
// Take two integers from the user and find the difference

package main

import "fmt"

func main() {
	var largeNum int // Initalize first number variable
	var smallNum int // Initalize second number variable

	fmt.Print("Please input a large number: ")
	fmt.Scan(&largeNum) // Take user's first number input

	fmt.Print("Please input a number smaller than the previous: ")
	fmt.Scan(&smallNum) // Take user's second number input

	fmt.Println("I shall now calculate:", largeNum, "%", smallNum, "=", largeNum%smallNum) // Output result
}
