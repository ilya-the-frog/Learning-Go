package main

import "fmt"
import "time"

func main() {
	var publisher, writer, artist, title string
	var year, pageNumber int
	var grade float32

	publisher = "DizzyBooks Publishing Inc."
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	title = "Mr. GoToSleep"
	year = 1997
	pageNumber = 14
	grade = 6.5
	//shantRate is my somehow created rating
	shantRate := float32(time.Now().Year()-1997) * grade

	fmt.Println("Comic:", title, "was written by", writer, "drawn by", artist)
	fmt.Println("Is published in", year, "by", publisher, "and have", pageNumber,
		"pages.", "Our grade:", grade)
	fmt.Println("My rating:", shantRate)

	fmt.Println()
	fmt.Println()

	// publisher = "DizzyBooks Publishing Inc." we don’t have to reassign it here
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	title = "Epic Vol. 1"
	year = 2013
	pageNumber = 160
	grade = 9.0

	shantRate = float32(time.Now().Year()-1997) * grade

	fmt.Println("Comic:", title, "was written by", writer, "drawn by", artist)
	fmt.Println("Is published in", year, "by", publisher, "and have", pageNumber,
		"pages.", "Our grade:", grade)
	fmt.Println("My rating:", shantRate)
}
