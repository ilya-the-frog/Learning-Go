package main

import "fmt"

func main () {
  fmt.Println("Put your's salary gross")
  var salary int
  fmt.Scan(&salary)
  netto := float32(salary) * 0.87

  fmt.Println("")
  fmt.Printf("Your salary netto is %.2f", netto)

}
