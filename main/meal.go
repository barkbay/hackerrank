package main

import (
	"os"

	"morello.ovh/hackerrank/30_days_of_code/operators"
)

func main() {
	// Read the matrix from Stdin
	m := operators.NewMeal(os.Stdin)
	m.Print()
}
