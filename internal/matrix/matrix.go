// Package matrix supplies matrix types and operations
package matrix

import (
	"fmt"
	"strings"
)

// Matrix is a basic implementation of a matrix.
type Matrix struct {
	rows, cols int
	data       []float64
}

// Get returns the float64 value at the given row and column
func (m Matrix) Get(row, col int) float64 {
	return m.data[m.cols*row+col]
}

// Dimensions returns the number of rows and columns a matrix has
func (m Matrix) Dimensions() (rows, cols int) {
	return m.rows, m.cols
}

func (m Matrix) set(row, col int, value float64) {
	m.data[m.cols*row+col] = value
}

func (m Matrix) String() string {
	writeRow := func(row []float64, sb *strings.Builder) {
		for i, v := range row {
			sb.WriteString(fmt.Sprint(v))
			if i != len(row)-1 {
				sb.WriteRune(' ')
			}
		}
	}

	// TODO: Implement this in terms of Get()
	sb := strings.Builder{}
	for len(m.data) > 0 {
		writeRow(m.data[0:m.cols], &sb)
		m.data = m.data[m.cols:]
		sb.WriteRune('\n')
	}
	return sb.String()[0 : sb.Len()-1]
}

func dimCheck(rows, cols int) {
	if rows <= 0 || cols <= 0 {
		panic("matrix: rows and cols cannot be less than or equal to 0")
	}
}

// New returns a matrix with all values set to 0
//
// Will panic if rows or cols is less than or equal to 0
func New(rows, cols int) Matrix {
	dimCheck(rows, cols)
	b := Matrix{rows: rows, cols: cols}
	b.data = make([]float64, rows*cols)
	return b
}

// NewFromSlice returns a matrix with all values imported from the
// supplied slice
func NewFromSlice(data []float64, rows, cols int) Matrix {
	dimCheck(rows, cols)
	if len(data) != rows*cols {
		err := fmt.Errorf(
			"matrix: supplied slice (%T) is expected to have a length of %d, instead its length is %d",
			data,
			rows*cols,
			len(data),
		)
		panic(err)
	}
	b := Matrix{rows: rows, cols: cols, data: data}
	return b
}

// Scale scales all of the entries in a matrix by multiplying them with the
// provided scalar, and returns a new matrix with the result.
func Scale(mat Matrix, scalar float64) Matrix {
	rows, cols := mat.Dimensions()
	result := New(rows, cols)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			result.set(r, c, mat.Get(r, c)*scalar)
		}
	}
	return result
}

// Add adds two matrices together and returns the result.
func Add(first Matrix, second Matrix) Matrix {
	rows, cols := first.Dimensions()
	r, c := second.Dimensions()
	if rows != r || cols != c {
		panic("matrix: the dimensions of the supplied matrices must be exactly equal.")
	}
	result := New(rows, cols)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			result.set(r, c, first.Get(r, c)+second.Get(r, c))
		}
	}
	return result
}
