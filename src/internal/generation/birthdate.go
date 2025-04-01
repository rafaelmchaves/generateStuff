package generation

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateBirthdate() string {
	maxAge := 110

	rand.Seed(time.Now().UnixNano())
	randomAge := rand.Intn(maxAge)

	now := time.Now()

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

// Generate birthdate thinking in slots and odds(e.g should generate less people who is more than 80 years old)
func GenerateBirthdateUsingOddsSlot(ageSlots []*AgeSlot) string {
	maxAge := 110

	rand.Seed(time.Now().UnixNano())
	tries := 0
	selectedSlot := -1
	randomAge := -1
	for selectedSlot == -1 && tries < 50 {
		randomAge = rand.Intn(maxAge)
		for i, slot := range ageSlots {
			if randomAge >= slot.MinAge && randomAge < slot.MaxAge && slot.LeftQuantity > 0 {
				selectedSlot = i
				slot.LeftQuantity -= 1
				tries = 0
				break
			}

		}

		tries++
	}

	if selectedSlot != -1 {
		return generateBirthdateByAge(randomAge)
	}

	fmt.Println("no one to be created anymore")
	return ""
}

func generateBirthdateByAge(age int) string {
	rand.Seed(time.Now().UnixNano())

	now := time.Now()
	birthYear := now.Year() - age

	month := time.Month(rand.Intn(12) + 1)

	dayOptions := daysInMonth(month, birthYear)
	day := rand.Intn(dayOptions) + 1

	return fmt.Sprintf("%02d/%02d/%d", day, time.Month(month), birthYear)
}

type AgeSlot struct {
	MinAge       int
	MaxAge       int
	LeftQuantity int64 // quantity of people in this slot left to be created
}
