package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	DayOfBirth   int
	MonthOfBirth int
	YearOfBirth  int
}

type Profile struct {
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		FullName: FullName{
			FirstName: "Nataliia",
			LastName:  "Kakhnych",
		},
		BirthDate: BirthDate{
			DayOfBirth:   7,
			MonthOfBirth: 9,
			YearOfBirth:  2005,
		},
		NumberOfSiblings: 1,
		ZodiacSign:       '\u264D',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
