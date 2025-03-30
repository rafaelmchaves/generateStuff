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

		// Split the line into city name and population
		parts := strings.Split(line, "@")
		if len(parts) != 2 {
			continue // Skip inva
			// lid lines
		}

		population, err := strconv.Atoi(parts[1])
		if err != nil {
			continue // Skip if conversion fails
		}

		cities = append(cities, City{Name: parts[0], Population: population})
	}

	return cities, nil
}
