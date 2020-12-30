package matrix_test

import (
	"fmt"
	"testing"

	"example.com/m/internal/matrix"
)

func TestBasicString(t *testing.T) {
	tests := []struct {
		rows, cols int
		expected   string
	}{
		{
			1, 1, "0",
		},
		{
			3, 1, "0\n0\n0",
		},
		{
			3, 3, "0 0 0\n0 0 0\n0 0 0",
		},
		{
			5, 5, "0 0 0 0 0\n0 0 0 0 0\n0 0 0 0 0\n0 0 0 0 0\n0 0 0 0 ",
		},
	}
	for _, test := range tests {
		m := matrix.NewBasic(test.rows, test.cols)
		mString := fmt.Sprint(m)
		if mString != test.expected {
			t.Errorf(
				"expected a %dx%d matrix to produce the following string:\n\n"+
					"%s\n\ninstead it produced:\n\n%s",
				test.rows,
				test.cols,
				mString,
				test.expected,
			)
		}
	}
}

func TestNewBasicFromSlice1D() {
	tests := []struct {
		rows, cols int
		data       []float64
	}{
		{
			1, 1, []float64{15},
		},
		{
			3, 1, []float64{15, 64, 32},
		},
		{
			3, 3, []float64{1, 3, 9, 2, 4, 6, 7, 14, 21},
		},
	}
	for _, test := range tests {
		m := matrix.NewBasicFromSlice(test.data, test.rows, test.cols)
		for i, v := range test.data {
			if m.Get(i/row, i%row)
		}
	}
}
