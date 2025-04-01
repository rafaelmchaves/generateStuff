package main

import (
	"fmt"

	"github.com/rafaelmchaves/generatestuff/src/internal/generation"
)

func main() {
	// println("----------- Generate stuff - application ----------")

	// cities, err := filemanager.LoadCitiesFromFile()
	// if err != nil {
	// 	fmt.Println("error to load cities from file", err)
	// }

	// for _, city := range cities {
	// 	fmt.Println("City:", city.Name, "Population:", city.Population)
	// }

	// birthdate := generation.GenerateBirthdate()
	// println("birthdate", birthdate)

	// ageSlots := []*generation.AgeSlot{
	// 	{MinAge: 0, MaxAge: 4, LeftQuantity: 500},
	// 	{MinAge: 5, MaxAge: 9, LeftQuantity: 400},
	// 	{MinAge: 10, MaxAge: 14, LeftQuantity: 600},
	// 	{MinAge: 15, MaxAge: 19, LeftQuantity: 800},
	// 	{MinAge: 20, MaxAge: 24, LeftQuantity: 900},
	// 	{MinAge: 25, MaxAge: 29, LeftQuantity: 1000},
	// 	{MinAge: 30, MaxAge: 34, LeftQuantity: 1100},
	// 	{MinAge: 35, MaxAge: 39, LeftQuantity: 1200},
	// 	{MinAge: 40, MaxAge: 44, LeftQuantity: 900},
	// 	{MinAge: 45, MaxAge: 49, LeftQuantity: 700},
	// 	{MinAge: 50, MaxAge: 54, LeftQuantity: 500},
	// 	{MinAge: 55, MaxAge: 59, LeftQuantity: 300},
	// 	{MinAge: 60, MaxAge: 64, LeftQuantity: 200},
	// 	{MinAge: 65, MaxAge: 69, LeftQuantity: 100},
	// 	{MinAge: 70, MaxAge: 74, LeftQuantity: 50},
	// 	{MinAge: 75, MaxAge: 79, LeftQuantity: 20},
	// 	{MinAge: 80, MaxAge: 84, LeftQuantity: 10},
	// 	{MinAge: 85, MaxAge: 89, LeftQuantity: 5},
	// 	{MinAge: 90, MaxAge: 94, LeftQuantity: 2},
	// 	{MinAge: 95, MaxAge: 100, LeftQuantity: 1},
	// 	{MinAge: 100, MaxAge: 110, LeftQuantity: 1},
	// }

	// people := 1
	// start := time.Now()
	// for people <= 5000 {
	// 	birthdate := generation.GenerateBirthdateUsingOddsSlot(ageSlots)
	// 	if birthdate == "" {
	// 		break
	// 	}
	// 	people++
	// }
	// elapsed := time.Since(start)

	// fmt.Println("generation time executed in ", elapsed)
	// for _, slot := range ageSlots {
	// 	fmt.Printf("MinAge: %d, MaxAge: %d, LeftQuantity: %d\n", slot.MinAge, slot.MaxAge, slot.LeftQuantity)
	// }

	// println(generation.GenerateFirstName())

	countMap := make(map[string]int)

	i := 0
	for i < 4000 {
		name := generation.GenerateFirstNameWithProbabilities()
		if count, exists := countMap[name]; exists {
			countMap[name] = count + 1
		} else {
			countMap[name] = 0
		}

		i++

	}

	for name, count := range countMap {
		println(fmt.Sprintf("Name:%s Count: %d", name, count))
	}

	// nameList := filemanager.LoadNames()
	// for _, name := range nameList {
	// 	println(fmt.Sprintf("Name:%s Probability: %d", name.Description, name.Probability))
	// }

}
