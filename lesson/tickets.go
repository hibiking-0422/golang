package main

import (
	"fmt"
	"math/rand"
)

func main() {

	const (
		distance      = 62100000
		secondsPerDay = 86400
	)

	fmt.Printf("%-16v %4v %-10v %4v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("-------------------------------------------")

	for count := 0; count < 10; count++ {
		spaceline_name := ""
		trip_type := ""
		rocket_speed := rand.Intn(15) + 16
		price := 20.0 + rocket_speed
		days := distance / rocket_speed / secondsPerDay

		switch num := rand.Intn(3); num {
		case 0:
			spaceline_name = "SpaceX"
		case 1:
			spaceline_name = "Space Adventures"
		case 2:
			spaceline_name = "Virgin Galactic"
		}

		if num := rand.Intn(2); num != 0 {
			trip_type = "Round-trip"
			price *= 2
		} else {
			trip_type = "One-way"
		}

		fmt.Printf("%-16v %4v %-10v $%4v\n", spaceline_name, days, trip_type, price)
	}

}
