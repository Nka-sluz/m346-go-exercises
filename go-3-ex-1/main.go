package main

import "fmt"

const (
	Aries       = '\u2648' // Widder
	Taurus      = '\u2649' // Stier
	Gemini      = '\u264a' // Zwillinge
	Cancer      = '\u264b' // Krebs
	Leo         = '\u264c' // Löwe
	Virgo       = '\u264d' // Jungfrau
	Libra       = '\u264e' // Waage
	Scorpius    = '\u264f' // Skorpion
	Sagittarius = '\u2650' // Schütze
	Capricornus = '\u2651' // Steinbock
	Aquarius    = '\u2652' // Wassermann
	Pisces      = '\u2653' // Fische
)

func outputWithZodiacSign(p Person) {
	var zodiacSign rune = '?'

	if p.BirthDate.Month == 1 {
		if p.BirthDate.Day < 20 {
			zodiacSign = Capricornus
		} else {
			zodiacSign = Aquarius
		}
	}

	if p.BirthDate.Month == 2 {
		if p.BirthDate.Day < 19 {
			zodiacSign = Aquarius
		} else {
			zodiacSign = Pisces
		}
	}

	if p.BirthDate.Month == 3 {
		if p.BirthDate.Day < 21 {
			zodiacSign = Pisces
		} else {
			zodiacSign = Aries
		}
	}

	if p.BirthDate.Month == 4 {
		if p.BirthDate.Day < 21 {
			zodiacSign = Aries
		} else {
			zodiacSign = Taurus
		}
	}

	if p.BirthDate.Month == 5 {
		if p.BirthDate.Day < 22 {
			zodiacSign = Taurus
		} else {
			zodiacSign = Gemini
		}
	}

	if p.BirthDate.Month == 6 {
		if p.BirthDate.Day < 22 {
			zodiacSign = Gemini
		} else {
			zodiacSign = Cancer
		}
	}

	if p.BirthDate.Month == 7 {
		if p.BirthDate.Day < 23 {
			zodiacSign = Cancer
		} else {
			zodiacSign = Leo
		}
	}

	if p.BirthDate.Month == 8 {
		if p.BirthDate.Day < 23 {
			zodiacSign = Leo
		} else {
			zodiacSign = Virgo
		}
	}

	if p.BirthDate.Month == 9 {
		if p.BirthDate.Day < 23 {
			zodiacSign = Virgo
		} else {
			zodiacSign = Libra
		}
	}

	if p.BirthDate.Month == 10 {
		if p.BirthDate.Day < 23 {
			zodiacSign = Libra
		} else {
			zodiacSign = Scorpius
		}
	}

	if p.BirthDate.Month == 11 {
		if p.BirthDate.Day < 23 {
			zodiacSign = Scorpius
		} else {
			zodiacSign = Sagittarius
		}
	}

	if p.BirthDate.Month == 12 {
		if p.BirthDate.Day < 21 {
			zodiacSign = Sagittarius
		} else {
			zodiacSign = Capricornus
		}
	}
	// NOTE: The runes are defined above as constants.

	fmt.Printf("%s %s, born on %02d.%02d.%04d, has the zodiac sign %c.\n",
		p.FirstName, p.LastName, p.Day, p.Month, p.Year, zodiacSign)
}

type FullName struct {
	FirstName string
	LastName  string
}
type BirthDate struct {
	Day   byte
	Month byte
	Year  uint16
}

type Person struct {
	FullName
	BirthDate
}

func main() {
	grace := Person{FullName{"Grace", "Hopper"}, BirthDate{9, 12, 1906}}
	dennis := Person{FullName{"Dennis", "Ritchie"}, BirthDate{9, 9, 1941}}
	rick := Person{FullName{"Rick", "Astley"}, BirthDate{6, 2, 1966}}
	edsger := Person{FullName{"Edsger", "Dijkstra"}, BirthDate{11, 5, 1930}}
	alan := Person{FullName{"Alan", "Turing"}, BirthDate{23, 6, 1912}}

	outputWithZodiacSign(grace)
	outputWithZodiacSign(dennis)
	outputWithZodiacSign(rick)
	outputWithZodiacSign(edsger)
	outputWithZodiacSign(alan)
}
