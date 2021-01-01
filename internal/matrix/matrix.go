// Package matrix supplies matrix types and operations
package matrix

import (
	"fmt"
	"math/rand"
	"strings"
)

const seed = 0

var (
	random = rand.New(rand.NewSource(seed))
)

// Matrix is a basic implementation of a matrix.
type Matrix struct {
	rows, cols int
	data       []float64
	transpose  *Matrix
}

func (m *Matrix) accessCheck(row, col int) {
	if row < 0 || col < 0 {
		err := fmt.Errorf(
			"matrix: row and col (%d, %d) cannot be less than 0",
			row,
			col,
		)
		panic(err)
	}
	if row >= m.rows || col >= m.cols {
		err := fmt.Errorf(
			"matrix: row and col (%d, %d) "+
				"cannot be greater than or equal to the matrix's dimensions (%dx%d)",
			row,
			col,
			m.rows,
			m.cols,
		)
		panic(err)
	}
}

// Get returns the float64 value at the given row and column
func (m *Matrix) Get(row, col int) float64 {
	m.accessCheck(row, col)
	return m.data[m.cols*row+col]
}

// Dimensions returns the number of rows and columns a matrix has
func (m *Matrix) Dimensions() (rows, cols int) {
	return m.rows, m.cols
}

func (m *Matrix) set(row, col int, value float64) {
	m.accessCheck(row, col)
	m.data[m.cols*row+col] = value
}

func (m *Matrix) String() string {
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

// Transpose generates a Matrix's transpose if it has not already been generated
func (m *Matrix) Transpose() *Matrix {
	if m.transpose != nil {
		return m.transpose
	}

	m.transpose = New(m.cols, m.rows)
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			m.transpose.set(c, r, m.Get(r, c))
		}
	}
	m.transpose.transpose = m
	return m.transpose
}

func dimCheck(rows, cols int) {
	if rows <= 0 || cols <= 0 {
		panic("matrix: rows and cols cannot be less than or equal to 0")
	}
}

// New returns a matrix with all values set to 0
//
// Will panic if rows or cols is less than or equal to 0
func New(rows, cols int) *Matrix {
	dimCheck(rows, cols)
	b := &Matrix{rows: rows, cols: cols}
	b.data = make([]float64, rows*cols)
	return b
}

// NewFromSlice returns a matrix with all values imported from the
// supplied slice
func NewFromSlice(data []float64, rows, cols int) *Matrix {
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
	b := &Matrix{rows: rows, cols: cols, data: data}
	return b
}

// NewRandomNormal returns a matrix with all values sourced from Go's math/rand.NormFloat64
func NewRandomNormal(rows, cols int) *Matrix {
	// TODO: Find a way to test this?
	data := make([]float64, rows*cols)
	for i := range data {
		data[i] = random.NormFloat64()
	}
	return NewFromSlice(data, rows, cols)
}

// Scale scales all of the entries in a matrix by multiplying them with the
// provided scalar, and returns a new matrix with the result.
func Scale(mat *Matrix, scalar float64) *Matrix {
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
func Add(first *Matrix, second *Matrix) *Matrix {
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

// Multiply multiplies two matrices together and returns the result.
func Multiply(first *Matrix, second *Matrix) *Matrix {
	if first.cols != second.rows {
		err := fmt.Errorf(
			"matrix: the cols of the first matrix (%dx%d) "+
				"must be equal to the rows of the second matrix (%dx%d)",
			first.rows, first.cols, second.rows, second.cols,
		)
		panic(err)
	}

	result := New(first.rows, second.cols)
	rows, cols := result.Dimensions()
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			var sum float64 = 0
			for offset := 0; offset < first.cols; offset++ {
				sum += first.Get(row, offset) * second.Get(offset, col)
			}
			result.set(row, col, sum)
		}
	}
	return result
}

// Map runs the given function on every entry in the matrix and returns the result
func Map(mat *Matrix, function func(float64) float64) *Matrix {
	rows, cols := mat.Dimensions()
	result := New(rows, cols)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			result.set(r, c, function(mat.Get(r, c)))
		}
	}
	return result
}
