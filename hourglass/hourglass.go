package hourglass

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// Matrix is basivally a 3x3 array extracted from a matrix
type Matrix struct {
	array [6][6]uint64
}

// Hourglass is basivally a 6x6 array
type Hourglass struct {
	array [3][3]uint64
}

// NewHourglass extracts an 3x3 hourglass from a Matrix
func (m *Matrix) NewHourglass(x, y uint64) Hourglass {
	result := [3][3]uint64{}
	// Extract the 3 lines
	subarray := m.array[x : x+3]
	// line 1
	result[0][0] = subarray[0][y]
	result[0][1] = subarray[0][y+1]
	result[0][2] = subarray[0][y+2]
	// line 2
	result[1][0] = subarray[1][y]
	result[1][1] = subarray[1][y+1]
	result[1][2] = subarray[1][y+2]
	// line 3
	result[2][0] = subarray[2][y]
	result[2][1] = subarray[2][y+1]
	result[2][2] = subarray[2][y+2]

	return Hourglass{result}
}

// NewMatrix creates from input stream
// 6 lines with unsigned int delemited by spaces are expected
func NewMatrix(in io.Reader) Matrix {
	if in == nil {
		in = os.Stdin
	}

	scanner := bufio.NewReader(in)
	array := [6][6]uint64{}

	for i := 0; i < 6; i++ {
		s, _ := scanner.ReadString('\n')
		ass := strings.Split(strings.TrimRight(s, "\n"), " ")
		j := 0
		for _, v := range ass {
			value, _ := strconv.ParseUint(v, 10, 64)
			array[i][j] = uint64(value)
			j++
		}
	}

	return Matrix{array}

}

func main() {
	NewMatrix(nil)
}
