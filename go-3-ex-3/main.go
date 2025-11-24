package main

import "fmt"

const (
	Lower = 1
	Upper = 30
)

func main() {
	for i := Lower; i <= Upper; i++ {
		toPrint := ""

		switch {
		case i%3 == 0:
			toPrint += "Fizz"
			if i%5 != 0 { // prevent unwanted Buzz
				fmt.Println(toPrint)
				break
			}
			fallthrough

		case i%5 == 0:
			toPrint += "Buzz"
			fmt.Println(toPrint)

		default:
			fmt.Println(i)
		}
	}

}
