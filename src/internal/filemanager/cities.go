package filemanager

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type City struct {
	Name       string
	Population int
}

func LoadCitiesFromFile() ([]City, error) {
	wd, _ := os.Getwd()
	fmt.Println("Current Working Directory:", wd)
	file, err := os.Open("/Users/rafael/programming projects/generateStuff/data/cities.txt")
	if err != nil {
		panic(err)
	}

	var cities []City
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		city := extractCityFromLine(line)

		cities = append(cities, city)
	}

	return cities, nil
}

func extractCityFromLine(line string) City {
	parts := strings.Split(line, "@")
	if len(parts) != 2 {
		panic("cities file is in the wrong format")
	}

	population, err := strconv.Atoi(parts[1])
	if err != nil {
		panic("format error. It was not possible to convert population to integer")
	}

	return City{Name: parts[0], Population: population}
}
