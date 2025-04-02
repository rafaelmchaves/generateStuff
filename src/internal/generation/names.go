package generation

import (
	"math/rand"
	"sort"
	"time"

	"github.com/rafaelmchaves/generatestuff/src/internal/filemanager"
)

func GenerateFirstName() string {

	nameList := filemanager.LoadNames()

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(len(nameList))

	return nameList[randomNumber].Description
}

// Names with more probabilities will have more chances to be generated.
func GenerateFirstNameWithProbabilities() string {
	nameList := filemanager.LoadNames()

	return generateWithProbability(nameList)
}

func GenerateSurnameNameWithProbabilities() string {
	nameList := filemanager.LoadEnglishSurnames()

	return generateWithProbability(nameList)
}

func GenerateSurname() string {

	nameList := filemanager.LoadEnglishSurnames()

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(len(nameList))

	return nameList[randomNumber].Description
}

func generateWithProbability(nameList []filemanager.Name) string {
	rand.Seed(time.Now().UnixNano())
	//generate a percentage(salt) in order to modify the probabilities
	quantityNamesToBeUsed := addProbabilities(nameList)

	if quantityNamesToBeUsed == 0 {
		return nameList[0].Description
	}

	//Sort name list by Probability in descending order
	sort.Slice(nameList, func(i, j int) bool {
		return nameList[i].Probability >= nameList[j].Probability
	})

	newNameList := insertBiggestOddsInTheMiddle(nameList[0:quantityNamesToBeUsed])
	randomNamePosition := rand.Intn(quantityNamesToBeUsed)

	return newNameList[randomNamePosition].Description
}

func GenerateFullName() string {
	return GenerateFirstName() + " " + GenerateSurname() + " " + GenerateSurname()
}

func GenerateFullNameWithProbabilities() string {
	return GenerateFirstNameWithProbabilities() + " " + GenerateSurnameNameWithProbabilities() + " " + GenerateSurnameNameWithProbabilities()
}

func addProbabilities(nameList []filemanager.Name) int {
	salt := rand.Intn(100)
	if salt > 80 {
		salt = 100
	}
	percentage := float64(salt) / 100.00

	return int(float64(len(nameList)) * percentage)
}

// The rand algorithm gives the middles position more probability to be generated.
// So, it was included the first position names to the middle in order to improve their probabilities
func insertBiggestOddsInTheMiddle(nameList []filemanager.Name) []filemanager.Name {
	size := len(nameList)
	if size > 15 {
		var i int
		for i <= 7 {
			aux := nameList[i]
			middle := size / 2
			nameList[i] = nameList[middle+i]
			nameList[middle] = aux
			i++
		}

	}
	return nameList
}
