/*
	Brian Padilla
	1/30/16
	Go language variable use with a Back to the Future theme.

	Note: All quotes and references belong to the original creators of the film.
*/

package main
import "fmt"

func main() {
	// Test variables
	test1 := "Are you telling me you built a time machine, out of a DeLorean?" // Marty McFly
	test2 := 88 // When you see this baby hit 88 miles per hour, you're gonna see some serious sh*t.
	test3 := 1.21 // 1.21 jigawatts! Great scott!
	test4 := true // November 5, 1955
	test5 := []string{"Great", "Scott", "!"} // Doc Brown

	// Outputs
	fmt.Printf("%T\n", test1) // prints string
	fmt.Printf("%T\n", test2) // prints int
	fmt.Printf("%T\n", test3) // prints float64
	fmt.Printf("%T\n", test4) // prints bool
	fmt.Printf("%T\n", test5) // prints []string
}
