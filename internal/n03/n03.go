package n03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var r1 = regexp.MustCompile(`[0-9\.]`)
var r2 = regexp.MustCompile(`[0-9]`)

type Matrix struct {
	Rows int
	Cols int
	data []string
}

func NewMatrix(rows, cols int) *Matrix {
	return &Matrix{
		Rows: rows,
		Cols: cols,
		data: make([]string, rows*cols),
	}
}

func (m *Matrix) At(row, col int) string {
	return m.data[row*m.Cols+col]
}

func (m *Matrix) Set(row, col int, v string) {
	m.data[row*m.Cols+col] = v
}

func (m *Matrix) Print() {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			print(m.At(i, j))
		}
		println()
	}
}

func GetMatrix(data string) *Matrix {
	rows := strings.Split(data, "\n")

	matrix := NewMatrix(len(rows), len(rows[0]))

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			matrix.Set(i, j, string(rows[i][j]))
		}
	}

	return matrix
}

func (m *Matrix) GetEngineSchematicSum() int {

	sum := 0

	for i := 0; i < m.Rows; i++ {

		for j := 0; j < m.Cols; j++ {
			if !r1.MatchString(m.At(i, j)) {
				// fmt.Printf("Found a symbol '%s' at [%d,%d]\n", m.At(i, j), i, j)

				// Scan the nearby area
				for k := i - 1; k <= i+1; k++ {
					for l := j - 1; l <= j+1; l++ {
						// Check to not go out of bounds
						if k >= 0 && k < m.Rows && l >= 0 && l < m.Cols {

							// If it's a number
							if r2.MatchString(m.At(k, l)) {
								numStr := getNumber(m, k, l)

								// fmt.Printf("Found a number '%s' at [%d,%d]\n", numStr, k, l)
								if num, err := strconv.Atoi(numStr); err == nil {
									sum += num
								}
							}
						}
					}
				}

			}
		}
	}

	return sum
}

func getNumber(m *Matrix, i, j int) string {

	// Out of bounds
	if j < 0 || j >= m.Cols {
		return ""
	}

	// Is not a number
	if !r2.MatchString(m.At(i, j)) {
		return ""
	}

	s := m.At(i, j)

	// Overwrite with a dot in order to no get the number twice
	m.Set(i, j, ".")

	return fmt.Sprintf("%s%s%s", getNumber(m, i, j-1), s, getNumber(m, i, j+1))
}
