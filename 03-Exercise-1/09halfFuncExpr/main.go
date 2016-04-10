// Brian Padilla
// 1/31/16
// Same program as 08Half.go, however this uses a function expression instead of a function declaration

package main

import "fmt"

func main() {
	half := func(n int) (int, bool) { // Half function expression
		return n/2, n%2 == 0 // Takes an integer and returns the number divided by two and a bool if it's even or odd
	}
	fmt.Println(half(1)) // Input and output first result
	fmt.Println(half(2)) // Input and output second result
}
