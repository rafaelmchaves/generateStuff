package internal

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateBirthdate() string {
	maxAge := 110

	rand.Seed(time.Now().UnixNano())
	randomAge := rand.Intn(maxAge)

	// Get the current time
	now := time.Now()

	// Calculate the birth year
	birthYear := now.Year() - randomAge

	month := time.Month(rand.Intn(12) + 1)

	dayOptions := daysInMonth(month, birthYear)
	day := rand.Intn(dayOptions) + 1

	return fmt.Sprintf("%02d/%02d/%d", day, time.Month(month), birthYear)
}

func daysInMonth(month time.Month, year int) int {
	switch month {
	case time.April, time.June, time.September, time.November:
		return 30
	case time.February:
		// Check if the year is a leap year
		if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
			return 29
		}
		return 28
	default:
		return 31
	}
}
