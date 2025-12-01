package main

import (
	"errors"
	"fmt"
	"os"
)

func computeGrade(gotPoints int, maxPoints int) (float32, error) {
	if gotPoints < 0 {
		return 0.0, errors.New("got less than 0 points")
	} else if gotPoints > maxPoints {
		return 0.0, errors.New("got more than maximum points")
	} else if maxPoints <= 0 {
		return 0.0, errors.New("max points is less or equal 0")
	}
	res := (float32(gotPoints)/float32(maxPoints))*5 + 1
	return res, nil
}

func main() {
	grade, err := computeGrade(1, 6)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Grade could not be calculated: %v\n", err)
	} else {
		fmt.Printf("Final grade is %v\n", grade)
	}
}
