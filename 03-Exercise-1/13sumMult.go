/* Brian Padilla
 * 3/21/16
 * Project Euler Problem 1: https://projecteuler.net/problem=1
 * If we list all the natural numbers below 10 that are multiples of 3 or 5,
 * we get 3, 5, 6 and 9. The sum of these multiples is 23.
 * Find the sum of all the multiples of 3 or 5 below 1000.
 */

package main

import (
  "fmt"
)

func sumMult(forRange int) {
  sum := 0
  for i := 0; i < forRange; i++ {
    // if the multiple is 3 or 5, then keep adding the sum of the previous
    if i%3 == 0 || i%5 == 0 {
      sum = sum + i // sum it up!
    }
  }
  // print the given range and the final result
  fmt.Println("The sum of all the multiples of 3 or 5 below",forRange,"is:",sum)
}

func main() {
  sumMult(10) // check the given example '10'
  sumMult(1000) // check the actual problem '1000'
}
