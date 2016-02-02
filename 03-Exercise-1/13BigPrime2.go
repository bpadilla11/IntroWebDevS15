// Brian Padilla
// 2/1/16
// 10001th Prime Number

package main

import "fmt"

func main() {
	var n int = 10
	var i int = 3
	//var count int
	//var c int

	for count := 2; count <= n; {
		for i := 2; i < 3; i++ {
			if i%3 == 0 {
				break
			}
			if (3 == i) {
				fmt.Println(i)
				count += 1
			}
		}
		i += 1
	}
}

/*
  This code is a work in progress and is subject to change or complete alteration.
*/
