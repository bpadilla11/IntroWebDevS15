// Brian Padilla
// 1/31/16
// Print different set of variables

package main

import "fmt"

// Function to take a set of integers
func foo(values ...int) {
	fmt.Println(values) // Print the set of integers
}

// Example given by the intructor
// Must write a function to make use of every instance (see above function)
func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
