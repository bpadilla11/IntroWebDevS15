// Brian Padilla
// 1/31/16
// Program that checks if a number is even or odd and the number divided by two

package main

import "fmt"

// Half function
// Takes an integer and returns the number divided by two and a bool if it's even or odd
func half(n int) (int, bool) {
	return n/2, n%2 == 0
}

func main() {
	a, odd := half(1) // Declare first test: a takes the input
	b, even := half(2) // Declare second test: b takes the input
	fmt.Println(a, odd) // Output first test
	fmt.Println(b, even) // Output second test
}
