package main

import (
  "fmt"
)

func main(){

  var a [] int
  a[2] = 20
  b := [20]int {1,3,4,56,6}
  var x int = 10
  t := 12
  var y int = 7
  var sum int = x + y + t

  if sum > 20 {
    fmt.Println("WADDUP")
  }

  fmt.Println(sum)
  fmt.Println(a)
  fmt.Println(b)
  b.append(b, 324)
  fmt.Println(b)
}
