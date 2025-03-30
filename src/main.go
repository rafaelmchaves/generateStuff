package main

import (
	"fmt"

	"github.com/rafaelmchaves/generatestuff/src/internal"
)

func main() {

	println("Generate stuff - application")
	birthdate := internal.GenerateBirthdate()
	println("birthdate", birthdate)

	ageSlots := []*internal.AgeSlot{
		{MinAge: 0, MaxAge: 4, LeftQuantity: 500},
		{MinAge: 5, MaxAge: 9, LeftQuantity: 400},
		{MinAge: 10, MaxAge: 14, LeftQuantity: 600},
		{MinAge: 15, MaxAge: 19, LeftQuantity: 800},
		{MinAge: 20, MaxAge: 24, LeftQuantity: 900},
		{MinAge: 25, MaxAge: 29, LeftQuantity: 1000},
		{MinAge: 30, MaxAge: 34, LeftQuantity: 1100},
		{MinAge: 35, MaxAge: 39, LeftQuantity: 1200},
		{MinAge: 40, MaxAge: 44, LeftQuantity: 900},
		{MinAge: 45, MaxAge: 49, LeftQuantity: 700},
		{MinAge: 50, MaxAge: 54, LeftQuantity: 500},
		{MinAge: 55, MaxAge: 59, LeftQuantity: 300},
		{MinAge: 60, MaxAge: 64, LeftQuantity: 200},
		{MinAge: 65, MaxAge: 69, LeftQuantity: 100},
		{MinAge: 70, MaxAge: 74, LeftQuantity: 50},
		{MinAge: 75, MaxAge: 79, LeftQuantity: 20},
		{MinAge: 80, MaxAge: 84, LeftQuantity: 10},
		{MinAge: 85, MaxAge: 89, LeftQuantity: 5},
		{MinAge: 90, MaxAge: 94, LeftQuantity: 2},
		{MinAge: 95, MaxAge: 100, LeftQuantity: 1},
	}

	people := 1
	for people <= 50000 {
		birthdate := internal.GenerateBirthdateUsingOddsSlot(ageSlots)
		fmt.Println("person:", people, "birthdate:", birthdate)
		if birthdate == "" {
			break
		}
		people++
	}

	for _, slot := range ageSlots {
		fmt.Printf("MinAge: %d, MaxAge: %d, LeftQuantity: %d\n", slot.MinAge, slot.MaxAge, slot.LeftQuantity)
	}

}
