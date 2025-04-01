package filemanager

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Name struct {
	Description string
	Probability int
}

func LoadNames() []Name {
	file, err := os.Open("/Users/rafael/programming projects/generateStuff/data/english_male_first_name.txt")
	if err != nil {
		panic(err)
	}

	var nameList []Name
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		name := extractNameFromLine(line)

		nameList = append(nameList, name)
	}

	return nameList
}

func extractNameFromLine(line string) Name {
	parts := strings.Split(line, "@")
	if len(parts) != 2 {
		print(line)
		panic("name file is in the wrong format ")
	}

	probability, err := strconv.Atoi(parts[1])
	if err != nil {
		panic("format error. It was not possible to convert probability to integer. Line:" + line + err.Error())
	}

	return Name{Description: parts[0], Probability: probability}
}
