// Brian Padilla
// 1/31/16
// FizzBuzz program

package main

import "fmt"

func main() {
	fmt.Println("Start FizzBuzz")
	for i := 0; i <= 100; i++ { // Loop 100 times
		if i%3 == 0 { // If number is a multiple of 3, output FIZZ
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 { // If number is a multiple of 5, output BUZZ
			fmt.Println(i, "Buzz")
		} else if i%3 == 0 && i%5 == 0 { // If number is a multiple of 3 and 5, output FizzBuzz
			fmt.Println(i, "FizzBuzz")
		} else { // Output remaining numbers
			fmt.Println(i)
		}
	} // End for loop
	fmt.Println("Process Complete!")
}
