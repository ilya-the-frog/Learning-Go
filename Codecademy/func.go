package main

import (
	"fmt"
	"time"
)

func joinTwoStrings(prefix string, suffix string) string {
	return prefix + suffix
}

func printOutDate() {
	t := time.Now()
	fmt.Println(t)
}

func waitForIt(message string) {
	defer fmt.Println("Done!")
	fmt.Println("Waiting")
	fmt.Println("Waiting")
	fmt.Println("Waiting")
	fmt.Println(message)
}

func main() {
	printOutDate()
	waitForIt(joinTwoStrings("Hi", " there"))
}
