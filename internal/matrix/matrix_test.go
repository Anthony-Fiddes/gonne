package matrix_test

import (
	"fmt"
	"testing"

	"example.com/m/internal/matrix"
)

func TestBasicString(t *testing.T) {
	tests := []struct {
		name       string
		rows, cols int
		expected   string
	}{
		{
			"1x1", 1, 1, "0",
		},
		{
			"3x1", 3, 1, "0\n0\n0",
		},
		{
			"3x3", 3, 3, "0 0 0\n0 0 0\n0 0 0",
		},
		{
			"5x5", 5, 5, "0 0 0 0 0\n0 0 0 0 0\n0 0 0 0 0\n0 0 0 0 0\n0 0 0 0 0",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := matrix.NewBasic(test.rows, test.cols)
			mString := fmt.Sprint(m)
			if mString != test.expected {
				t.Fatalf(
					"expected a %dx%d matrix to produce the following string:\n\n"+
						"%s\n\ninstead it produced:\n\n%s",
					test.rows,
					test.cols,
					mString,
					test.expected,
				)
			}
		})
	}
}

func TestNewBasicFromSlice(t *testing.T) {
	tests := []struct {
		name       string
		rows, cols int
		data       []float64
	}{
		{
			"1x1", 1, 1, []float64{15},
		},
		{
			"3x1 Vector", 3, 1, []float64{15, 64, 32},
		},
		{
			"3x3 Matrix", 3, 3, []float64{1, 3, 9, 2, 4, 6, 7, 14, 21},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := matrix.NewBasicFromSlice(test.data, test.rows, test.cols)
			for i, v := range test.data {
				row := i / test.rows
				col := i % test.rows
				if v != m.Get(row, col) {
					t.Fatalf(
						"At row %d, col %d the matrix was expected to return %f"+
							"as prescribed in the test data (%v). Instead it "+
							"returned %f",
						row,
						col,
						v,
						test.data,
						m.Get(row, col),
					)
				}
			}
		})
	}
}
