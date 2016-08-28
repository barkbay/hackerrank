package operators

import (
	"fmt"
	"strings"
	"testing"
)

func TestTotalPrice(t *testing.T) {
	f1 := strings.NewReader(`12.00
20
8`)
	m := NewMeal(f1)
	fmt.Printf("%f", m.GetTotalPrice())
}
