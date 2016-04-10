// Brian Padilla
// 1/31/16
// Output result of natural numbers

package main

import "fmt"

func main() {
	count := 0 // Start counter at 0
	for i := 0; i < 1000; i++ { // loop 1000 times
		if i%3 == 0 { // if number is 3, add to counter
			count += i
		} else if i%5 == 0 { // if number is 5, add to counter
			count += i
		}
	}
	fmt.Println(count) // print the sum of the counter
}
