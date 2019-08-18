package main

import (
  "fmt"
)

func main() {
  //setting a go routine with a buffer
  c := make(chan string, 2)
  c <- "hello"
  c <- "world"

  msg := <- c
  fmt.Println(msg)

  msg = <- c
  fmt.Println(msg)
}
