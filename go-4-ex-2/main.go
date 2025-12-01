package main

import (
	"fmt"
	"math"
)

type ShortSides struct {
	a float64
	b float64
}

func computeHypotenuse(a float64, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func (shortSides ShortSides) Hypotenuse() float64 {
	return math.Sqrt(math.Pow(shortSides.a, 2) + math.Pow(shortSides.b, 2))
}

func main() {
	y := computeHypotenuse(3, 4)
	fmt.Println("Computed hypotenuse: ", y)                    // res - 5
	fmt.Println("Hypotenuse: ", ShortSides{3, 4}.Hypotenuse()) // res - 5
}
