package hourglass

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	f1 := strings.NewReader(
		`-1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 2 4 4 0
0 0 0 2 0 0
0 0 1 2 4 0`)
	actual := NewMatrix(f1)

	// print matrix in this unit test
	for _, i := range actual.array {
		for _, j := range i {
			fmt.Printf("%d ", j)
		}
		fmt.Println("")
	}

	expectedMatrix := [6][6]int64{
		[6]int64{-1, 1, 1, 0, 0, 0},
		[6]int64{0, 1, 0, 0, 0, 0},
		[6]int64{1, 1, 1, 0, 0, 0},
		[6]int64{0, 0, 2, 4, 4, 0},
		[6]int64{0, 0, 0, 2, 0, 0},
		[6]int64{0, 0, 1, 2, 4, 0}}

	expected := Matrix{expectedMatrix}
	if !reflect.DeepEqual(actual, expected) {
		t.Fail()
	}
}

func TestNewHourglass1(t *testing.T) {
	f1 := strings.NewReader(
		`1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 2 4 4 0
0 0 0 2 0 0
0 0 1 2 4 0`)
	m := NewMatrix(f1)
	expectedMatrix1 := [3][3]int64{
		[3]int64{1, 1, 1},
		[3]int64{0, 1, 0},
		[3]int64{1, 1, 1}}
	actual := m.NewHourglass(0, 0).array
	if !reflect.DeepEqual(actual, expectedMatrix1) {
		t.Fail()
	}
}

func TestNewHourglass2(t *testing.T) {
	f1 := strings.NewReader(
		`1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 2 4 4 0
0 0 0 2 0 0
0 0 1 2 4 0`)
	m := NewMatrix(f1)
	expectedMatrix1 := [3][3]int64{
		[3]int64{1, 0, 0},
		[3]int64{2, 4, 4},
		[3]int64{0, 2, 0}}
	actual := m.NewHourglass(2, 2).array
	if !reflect.DeepEqual(actual, expectedMatrix1) {
		t.Fail()
	}
}

func TestHourglassSum(t *testing.T) {
	f1 := strings.NewReader(
		`1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 2 4 4 0
0 0 0 2 0 0
0 0 1 2 4 0`)
	m := NewMatrix(f1)
	h := m.NewHourglass(2, 2)
	if h.sum() != 7 {
		t.Fail()
	}
}

func TestMaxtrixSum(t *testing.T) {
	f1 := strings.NewReader(
		`1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 2 4 4 0
0 0 0 2 0 0
0 0 1 2 4 0`)
	m := NewMatrix(f1)
	if m.LargestSum() != 19 {
		t.Fail()
	}
}

func TestMaxtrixSum2(t *testing.T) {
	f1 := strings.NewReader(
		`1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 9 2 -4 -4 0
0 0 0 -2 0 0
0 0 -1 -2 -4 0`)
	m := NewMatrix(f1)
	if m.LargestSum() != 13 {
		t.Fail()
	}
}

func TestMaxtrixSum3(t *testing.T) {
	f1 := strings.NewReader(
		`1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 9 2 -4 4 0
0 0 0 -2 2 0
0 0 -1 -2 4 10`)
	m := NewMatrix(f1)
	if m.LargestSum() != 14 {
		t.Fail()
	}
}

func TestMaxtrixSum4(t *testing.T) {
	f1 := strings.NewReader(
		`-1 -1 0 -9 -2 -2
-2 -1 -6 -8 -2 -5
-1 -1 -1 -2 -3 -4
-1 -9 -2 -4 -4 -5
-7 -3 -3 -2 -9 -9
-1 -3 -1 -2 -4 -5`)
	m := NewMatrix(f1)
	if m.LargestSum() != -6 {
		t.Fail()
	}
}
