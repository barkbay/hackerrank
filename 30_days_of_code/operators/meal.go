package operators

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Meal represent a meal
type Meal struct {
	mealCost, tipPercent, taxPercent float64
}

// NewMeal build a new meal
func NewMeal(in io.Reader) Meal {
	var m Meal
	_, err := fmt.Scanf("%f", &m.mealCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Scanf("%f", &m.tipPercent)
	fmt.Scanf("%f", &m.taxPercent)
	return m
}

func (m *Meal) getTaxes() float64 {
	return (m.mealCost * m.taxPercent) / float64(100)
}

func (m *Meal) getTip() float64 {
	return (m.mealCost * m.tipPercent) / float64(100)
}

// GetTotalPrice computes the total price
func (m *Meal) GetTotalPrice() float64 {
	return m.mealCost + m.getTaxes() + m.getTip()
}

// Print print a sentence
func (m *Meal) Print() {
	fmt.Printf("The total meal cost is %.0f dollars.\n", m.GetTotalPrice())
}

func main() {
	// Read the matrix from Stdin
	m := NewMeal(os.Stdin)
	m.Print()
}
