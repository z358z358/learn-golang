// How do you specify the direction of a channel type?
// c chan<- receive-only , c <-chan send-only

package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("start")
  sleep(3)
  fmt.Println("end")
  // What is a buffered channel? How would you create one with a capacity of 20?
  // c := make(chan int, 20)
}

// Write your own Sleep function using time.After.
func sleep(sec int) {
  <-time.After(time.Second * time.Duration(sec))
}
