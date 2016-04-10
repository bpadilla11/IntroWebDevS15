// Brian Padilla
// 1/31/16
// Looks for the largest number is a set of numbers

package main

import "fmt"

// Function to take a multiple set of numbers
func biggest(numbers ...int) int {
	var largest int // Initalize largest variable
	for _, i := range numbers { // Loop through the given set
		if i > largest { // If the variable is larger, then it's the biggest
			largest = i // The bigger becomes the biggest
		}
	}
	return largest // Return the biggest variable found
}

func main() {
	greatest1 := biggest(17, 18, 4, 11, 10, 21, 1, 7, 13, 3) // First test set
	greatest2 := biggest(8, 7, 16, 18, 19, 9, 4, 14, 20, 1) // Second test set
	greatest3 := biggest(4, 9, 2, 13, 15, 14, 19, 11, 1, 17) // Third test set

	fmt.Println(greatest1) // 1 fish
	fmt.Println(greatest2) // 2 fish
	fmt.Println(greatest3) // 3 fish, Go!
}
