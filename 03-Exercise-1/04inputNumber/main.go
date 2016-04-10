// Brian Padilla
// 1/31/16
// Take two integers from the user and finda the difference remainder

package main

import "fmt"

func main() {
	var largeNum int // Initalize first number variable
	var smallNum int // Initalize second number variable

	fmt.Print("Hey buddy, what's the largest number you got?: ")
	fmt.Scan(&largeNum) // Take user's first number input

	fmt.Print("I'm not your buddy, friend. Now tell me a number smaller than your largest: ")
	fmt.Scan(&smallNum) // Take user's second number input

	fmt.Println("A'right guy, I'm aboot to calculate:", largeNum, "%", smallNum, "=", largeNum%smallNum) // Output result
}
