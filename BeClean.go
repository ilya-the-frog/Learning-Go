package main

import "fmt"

func main () {
  // solve https://thecode.media/ill-be-clean/
  // i didn't look at solution, done all myself
  var perimetr  int
  var height, square float

  height = 2.5
  perimetr = 50
  square = height * float(perimetr)

  var sanSquare int
  sanSquare = 5

  var sanNeeded float
  sanNeeded = square / sanSquare

  var famSize, sanPerDay, sanAvailable, cleanReg int
  famSize = 3
  sanPerDay = 4
  cleanReg = 2
  sanAvailable = famSize * sanPerDay * cleanReg

  if sanNeeded < sanAvailable {
      fmt.Println("House can be cleaned")
    }
    else {
      fmt.Println("House can't be cleaned")
    }
}
