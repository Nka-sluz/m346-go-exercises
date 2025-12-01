package main

import (
	"fmt"
	"math"
)

func computeQuadraticFormula(a float64, b float64, c float64) []float64 {
	d := math.Pow(b, 2) - 4*a*c
	if d < 0 {
		return []float64{}
	} else if d == 0 {
		return []float64{-b / (2 * a)}
	}
	x1 := (-b + math.Sqrt(d)) / (2 * a)
	x2 := (-b - math.Sqrt(d)) / (2 * a)
	return []float64{x1, x2}
}

func computeDiscriminant(a float64, b float64, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

func computeQuadraticFormulaWithComputeDiscriminant(a float64, b float64, c float64) []float64 {
	d := computeDiscriminant(a, b, c)
	if d < 0 {
		return []float64{}
	} else if d == 0 {
		return []float64{-b / (2 * a)}
	}
	x1 := (-b + math.Sqrt(d)) / (2 * a)
	x2 := (-b - math.Sqrt(d)) / (2 * a)
	return []float64{x1, x2}
}

func main() {
	fmt.Println(computeQuadraticFormula(3, 4, 1))                        // -0.33... -1
	fmt.Println(computeQuadraticFormula(2, 4, 2))                        // -1
	fmt.Println(computeQuadraticFormula(3, 4, 2))                        // leer
	fmt.Println(computeDiscriminant(3, 4, 1))                            // 4
	fmt.Println(computeDiscriminant(2, 4, 2))                            // 0
	fmt.Println(computeDiscriminant(3, 4, 2))                            // -8
	fmt.Println(computeQuadraticFormulaWithComputeDiscriminant(3, 4, 1)) // -0.33... -1
	fmt.Println(computeQuadraticFormulaWithComputeDiscriminant(2, 4, 2)) // -1
	fmt.Println(computeQuadraticFormulaWithComputeDiscriminant(3, 4, 2)) // leer
}
