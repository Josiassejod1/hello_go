package main

import (
  "fmt"
  "errors"
  "math"
)

func main() {
  vertices := make(map[string]int)
 result := sum(2,3)
  vertices["ashley"] = 23
  vertices["dalvin"] = 22

  fmt.Println(vertices["dalvin"])


  for i := 0; i < 5; i++ {
    fmt.Println(i)
  }

  for index, value := range vertices {
    fmt.Println("index:", index,"value", value)
  }

  fmt.Println(result)

  sqrt_result, err := sqrt(-5)

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Print(sqrt_result)
  }

}


func sum(x int, y int) int {
  return x + y
}

func sqrt(x  float64) (float64, error) {
  if x < 0 {
    return 0, errors.New("Undefined for negative value")
  } else {
    return math.Sqrt(x), nil
  }
}
