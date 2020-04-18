package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
  fmt.Println("You still have", fuel, "gallons")
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
  var fuel int
  switch planet {
  case "Venus":
    fuel = 300000
  case "Mercury":
    fuel = 500000
  case "Mars":
    fuel = 700000
  default:
    fmt.Println("We can't go here, no fuel used")
    fuel = 0
  }
  return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
  fmt.Println("Welcome to planet", planet)
}

// Create the function cantFly() here
func cantFly() {
  fmt.Println("We do not have the available fuel to fly.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
  var fuelRemaining, fuelCost int
  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)

  if fuel != 0 {
    if fuelRemaining >= fuelCost {
    greetPlanet(planet)
    fuelRemaining = fuelRemaining - fuelCost
    } else {
    cantFly()
    }
  } else {
      fmt.Println("Select again")
    }

  return fuelRemaining
}

func dest() {
  fmt.Println("Select destination")
}

func checkFuel(fuel int) int {
  if fuel >= 700000 {
    fmt.Println("We can go to Venus, Mercury or Mars")
    dest()
  } else if fuel >= 500000 {
    fmt.Println("We can go to Venus or Mercury")
    dest()
  } else if fuel >= 300000 {
    fmt.Println("We can go to Venus")
    dest()
  } else {
    cantFly()
  }
  return fuel
}

func main() {
  // Create `planetChoice` and `fuel`
  fmt.Println("Hello Space Traveller! How much fuel you have?")
  var fuel int
  fmt.Scan(&fuel)

  if fuel == 0 {
    fmt.Println("Try adding more fuel")
    fmt.Scan(&fuel)
  }

  var planetChoice string

  for fuel > 300000 {
    checkFuel(fuel)
    fmt.Scan(&planetChoice)
    fuel = flyToPlanet(planetChoice, fuel)
    fuelGauge(fuel)
  }
    cantFly()

}
