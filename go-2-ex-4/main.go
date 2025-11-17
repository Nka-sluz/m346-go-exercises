package main

import "fmt"

func main() {
	type Student struct {
		FirstName string
		LastName  string
	}

	type Class map[uint]Student

	classA := Class{
		1: {FirstName: "Ben", LastName: "Müller"},
		2: {FirstName: "Tim", LastName: "Muster"},
		3: {FirstName: "Tina", LastName: "Beispiel"},
	}

	classB := Class{
		1: {FirstName: "David", LastName: "Huber"},
		2: {FirstName: "Eva", LastName: "Kunz"},
		3: {FirstName: "Felix", LastName: "Meier"},
	}

	modules := map[uint][]Class{
		104: {classA},
		117: {classA, classB},
		346: {classB},
	}

	fmt.Println(modules)
}
