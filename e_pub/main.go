package main

import (
  "fmt"
)

func numGen(start, limit, skip int) <-chan int {
  out := make(chan int)

  go func() {
    for i := start; i < limit; {
      out <- i
      i = i + skip
    }
    close(out)
  }()
  return out
}


func main() {
 st := numGen(2, 1000, 2)
 for c := range st {
  fmt.Printf("%d, ", c)
 }
}
