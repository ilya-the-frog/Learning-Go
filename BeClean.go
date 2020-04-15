package main

import "fmt"

func main () {
  // solve https://thecode.media/ill-be-clean/
  // i didn't look at solution, done all myself
  var perimetr  int
  var height, square float32

  height = 2.5
  perimetr = 50
  square = height * float32(perimetr)

  var sanSquare int
  sanSquare = 5

  var sanNeeded float32
  sanNeeded = square / float32(sanSquare)

  var famSize, sanPerDay, cleanReg int
  var sanAvailable float32
  famSize = 3
  sanPerDay = 4
  cleanReg = 2
  sanAvailable = float32(famSize * sanPerDay * cleanReg)

  if sanNeeded <= sanAvailable {
      fmt.Println("House can be cleaned")
    } else {
      fmt.Println("House can't be cleaned")
    }
}
