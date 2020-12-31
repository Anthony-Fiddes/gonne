// Package matrix supplies matrix types and operations
package matrix

import (
	"fmt"
	"strings"
)

// Mat is a basic implementation of a matrix.
type Mat struct {
	rows, cols int
	data       []float64
}

// Get returns the float64 value at the given row and column
func (b Mat) Get(row, col int) float64 {
	return b.data[b.cols*row+col]
}

func (b Mat) set(row, col int, value float64) {
	b.data[b.cols*row+col] = value
}

func (b Mat) String() string {
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

// New returns a matrix with all values set to 0
//
// Will panic if rows or cols is less than or equal to 0
func New(rows, cols int) *Mat {
	dimCheck(rows, cols)
	b := &Mat{rows: rows, cols: cols}
	b.data = make([]float64, rows*cols)
	return b
}

// NewFromSlice returns a matrix with all values imported from the
// supplied slice
func NewFromSlice(data []float64, rows, cols int) *Mat {
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
	b := &Mat{rows: rows, cols: cols, data: data}
	return b
}
