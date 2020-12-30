// Package matrix supplies matrix types and operations
package matrix

import (
	"fmt"
	"strings"
)

// Matrix is an interface that describes the fundamental methods for doing
// matrix operations.
type Matrix interface {
	// Cols returns the number of columns
	Cols() int
	// Rows returns the number of rows
	Rows() int
	// Get returns the float64 value at the given row and column
	Get(row, col int) float64
}

// Transposable is an interface that describes a matrix which can supply its
// own transpose.
type Transposable interface {
	// T returns the transpose of the matrix
	T() Matrix
	Matrix
}

// Basic is a basic implementation of a matrix.
type Basic struct {
	cols int
	rows int
	data []float64
}

// Cols returns the number of columns
func (b Basic) Cols() int {
	return b.cols
}

// Rows returns the number of rows
func (b Basic) Rows() int {
	return b.rows
}

// Get returns the float64 value at the given row and column
func (b Basic) Get(row, col int) float64 {
	return b.data[b.cols*row+col]
}

func (b Basic) set(row, col int, value float64) {
	b.data[b.cols*row+col] = value
}

func (b Basic) String() string {
	writeRow := func(row []float64, sb *strings.Builder) {
		for i, v := range row {
			sb.WriteString(fmt.Sprint(v))
			if i != len(row)-1 {
				sb.WriteRune(' ')
			}
		}
	}

	sb := strings.Builder{}
	for len(b.data) > 0 {
		writeRow(b.data[0:b.cols], &sb)
		b.data = b.data[b.cols:]
		sb.WriteRune('\n')
	}
	return sb.String()[0 : sb.Len()-1]
}

func dimCheck(rows, cols int) {
	if rows <= 0 || cols <= 0 {
		panic("matrix: rows and cols cannot be less than or equal to 0")
	}
}

// NewBasic returns a Basic matrix with all values set to 0
//
// Will panic if rows or cols is less than or equal to 0
func NewBasic(rows, cols int) *Basic {
	dimCheck(rows, cols)
	b := &Basic{rows: rows, cols: cols}
	b.data = make([]float64, rows*cols)
	return b
}

// NewBasicFromSlice returns a Basic matrix with all values imported from the
// supplied slice
func NewBasicFromSlice(data []float64, rows, cols int) *Basic {
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
	b := &Basic{rows: rows, cols: cols, data: data}
	return b
}
