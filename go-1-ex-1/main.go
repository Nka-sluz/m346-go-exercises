package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!

	// Aufgabe
	var firstName string = "Nataliia"
	var lastName string = "Kakhnych"
	var dayOfBirth int = 7
	var monthOfBirth int = 9
	var yearOfBirth int = 2005
	var numberOfSiblings int = 1
	heightInMeters := 1.63
	var zodiacSign = 'J'
	// Aufgabe

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
