package main

import "fmt"

// solving https://thecode.media/gasoline/

func main() {
	var distance, oilPrice, oilUsage int
	var workDays, weeks int

	distance = 10
	oilPrice = 45
	oilUsage = 11 // per 100km!
	weeks = 4

	//count days at work
	workDays = weeks * 4

	//every day Andrew takes two rides so
	rides := workDays * 2

	totalDist := rides * distance

	var oilSpend float32
	oilSpend = float32(totalDist/100) * float32(oilUsage)

	var moneySpend float32
	moneySpend = oilSpend * float32(oilPrice)

	fmt.Println("Andrew would have spend on oil", moneySpend, "rub")

	var seasons, sPrice int
	seasons = 8
	sPrice = 300
	totalPrice := float32(seasons * sPrice)
	fmt.Println("All seasons costs", totalPrice, "rub")

	if totalPrice <= moneySpend {
		fmt.Println("Andrew can buy Game of the Thrones")
	} else {
		fmt.Println("Should add", totalPrice-moneySpend, "rub")
	}
}
