package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())

  var isHeistOn bool
  isHeistOn = true

  eludedGuards := rand.Intn(100)
  if eludedGuards > 50 {
    fmt.Println("Passed guards! Wooo")
  } else {
    isHeistOn = false
    fmt.Println("Plan better next time")
  }

  /* fmt.Println("")
  fmt.Println("Heist online?", isHeistOn)
  fmt.Println("") */

  openedVault := rand.Intn(100)
  if openedVault > 70 && isHeistOn == true {
    fmt.Println("Take EVERYTHING!!")
  } else {
    if isHeistOn == false {

    } else {
      isHeistOn = false
      fmt.Println("Vault closed. Police incoming....")
    }

  }

  /* fmt.Println("")
  fmt.Println("Heist online?", isHeistOn)
  fmt.Println("") */

  leftSafely := rand.Intn(5)
  if isHeistOn {
    switch leftSafely {
      case 0:
        isHeistOn = false
        fmt.Println("Alarm triggered, u lose")
      case 1:
        isHeistOn = false
        fmt.Println("Police spotted you, u lose")
      case 2:
        isHeistOn = false
        fmt.Println("Guards spotted you, u lose")
      case 3:
        isHeistOn = false
        fmt.Println("Visitors spotted you, u lose")
      case 4:
        fmt.Println("Start the getaway car!")
        amtStolen := 10000 + rand.Intn(1000000)
        fmt.Println("$", amtStolen, "not bad!")
    }
  } else {
    fmt.Println("U got caught")
  }

}
