package main

import (
  "fmt"
  "time"
)

func main() {
  go count("sheep")
      count("fish")
      fmt.Scanln()
}

func count(thing string) {
  for i:= 1; i <= 5; i++ {
    fmt.Println(i, thing)
    time.Sleep(time.Millisecond * 500)
  }
}
