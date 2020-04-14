package main

import "fmt"

func main () {
  fmt.Println("Put your's salary gross")
  var salary := scan()
  var netto float
  netto := salary * 0.87

  fmt.Println("Your salary netto is ", netto)
}
