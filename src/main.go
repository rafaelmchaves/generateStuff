package main

import "github.com/rafaelmchaves/generatestuff/src/internal"

func main() {

	println("Generate stuff - application")
	birthdate := internal.GenerateBirthdate()
	println("birthdate", birthdate)
}
