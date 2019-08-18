package main

import (
  "fmt"
)

type person struct {
  name string
  age int
  address string
}


func main(){
  p := person{name: "Dalvin", age: 22, address: "da block"}
  i := 7
  fmt.Println(&i)
  fmt.Println(p.name)
}
