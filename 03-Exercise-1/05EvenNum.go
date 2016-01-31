// Brian Padilla
// 1/31/16
// Print even numbers using a for loop

package main

import "fmt"

func main() {
	fmt.Println("Printing even numbers less than 100...")
	for i := 0; i <= 100; i++ { // Loop 100 times
		if i%2 == 0 { // If the number is even, then print it
			fmt.Println(i)
		}
	}
	fmt.Println("Process complete!")
}
