// Brian Padilla
// 1/31/16
// 10001th Prime Number

package main

import "fmt"

func main() {
  var flag bool
  var maxPrime int = 13

  for i := 2; i <= maxPrime; i++ {
    flag = false
    for j := 2; j < i; j++ {
      if (i%j==0) {
        flag = true
        break
      }
    }
    if flag == false {
      fmt.Println(i)
    }
  }
}

/*
  This code is a work in progress and is subject to change or complete alteration.
*/
