package matrix_test

import (
	"fmt"
	"testing"

	"example.com/m/internal/matrix"
)

func TestMatrixString(t *testing.T) {
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
			m := matrix.New(test.rows, test.cols)
			result := fmt.Sprint(m)
			if result != test.expected {
				t.Fatalf(
					"expected a %dx%d matrix to produce the following string:\n\n"+
						"%s\n\ninstead it produced:\n\n%s",
					test.rows,
					test.cols,
					result,
					test.expected,
				)
			}
		})
	}
}

func TestNewFromSlice(t *testing.T) {
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
			mat := matrix.NewFromSlice(test.data, test.rows, test.cols)
			index := 0
			for row := 0; row < test.rows; row++ {
				for col := 0; col < test.cols; col++ {
					result := mat.Get(row, col)
					expectedResult := test.data[index]
					if result != expectedResult {
						t.Fatalf(
							"At row %d, col %d the matrix was expected to return %f "+
								"as prescribed in the test data (%v). Instead it "+
								"returned %f",
							row,
							col,
							expectedResult,
							test.data,
							result,
						)
					}
					index++
				}
			}
		})
	}
}

func TestScale(t *testing.T) {
	tests := []struct {
		name       string
		rows, cols int
		data       []float64
		scalar     float64
	}{
		{
			"1x1", 1, 1, []float64{15}, 2,
		},
		{
			"3x1 Vector", 3, 1, []float64{15, 64, 32}, 3.15,
		},
		{
			"3x3 Matrix", 3, 3, []float64{1, 3, 9, 2, 4, 6, 7, 14, 21}, 1.28,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := matrix.NewFromSlice(test.data, test.rows, test.cols)
			scaled := matrix.Scale(m, test.scalar)
			index := 0
			for row := 0; row < test.rows; row++ {
				for col := 0; col < test.cols; col++ {
					expectedResult := test.data[index] * test.scalar
					result := scaled.Get(row, col)
					if result != expectedResult {
						t.Fatalf(
							"At row %d, col %d the matrix was expected to return %f "+
								"as prescribed in the test data (%v) multiplied by %f. "+
								"Instead it returned %f",
							row,
							col,
							expectedResult,
							test.data,
							test.scalar,
							result,
						)
					}
					index++
				}
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name       string
		rows, cols int
		first      []float64
		second     []float64
	}{
		{
			"1x1", 1, 1, []float64{15}, []float64{10},
		},
		{
			"3x1 Vector", 3, 1, []float64{15, 64, 32}, []float64{21, 32, 85},
		},
		{
			"3x3 Matrix", 3, 3, []float64{1, 3, 9, 2, 4, 6, 7, 14, 21}, []float64{0, 1, 0, 1, 0, 1, 0, 1, 185},
		},
		{
			"3x3 Matrix Identity", 3, 3, []float64{1, 3, 9, 2, 4, 6, 7, 14, 21}, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := matrix.NewFromSlice(test.first, test.rows, test.cols)
			addend := matrix.NewFromSlice(test.second, test.rows, test.cols)
			sum := matrix.Add(m, addend)
			index := 0
			for row := 0; row < test.rows; row++ {
				for col := 0; col < test.cols; col++ {
					expectedResult := test.first[index] + addend.Get(row, col)
					result := sum.Get(row, col)
					if result != expectedResult {
						t.Fatalf(
							"At row %d, col %d the matrix was expected to return %f "+
								"as prescribed in the test data (%v added to %v). "+
								"Instead it returned %f",
							row,
							col,
							expectedResult,
							test.first,
							test.second,
							result,
						)
					}
					index++
				}
			}
		})
	}
}

func TestMatrixTranspose(t *testing.T) {
	tests := []struct {
		name       string
		rows, cols int
		data       []float64
		expected   []float64
	}{
		{
			"1x1", 1, 1, []float64{15}, []float64{15},
		},
		{
			"3x1 Vector", 3, 1, []float64{15, 64, 32}, []float64{15, 64, 32},
		},
		{
			"3x3 Matrix", 3, 3, []float64{1, 3, 9, 2, 4, 6, 7, 14, 21}, []float64{1, 2, 7, 3, 4, 14, 9, 6, 21},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := matrix.NewFromSlice(test.data, test.rows, test.cols)
			expected := matrix.NewFromSlice(test.expected, test.cols, test.rows)
			transpose := m.Transpose()
			for row := 0; row < test.cols; row++ {
				for col := 0; col < test.rows; col++ {
					result := transpose.Get(row, col)
					expectedResult := expected.Get(row, col)
					if result != expectedResult {
						t.Fatalf(
							"At row %d, col %d the matrix was expected to return %f "+
								"as prescribed in the test data (%v). "+
								"Instead it returned %f",
							row,
							col,
							expectedResult,
							test.data,
							result,
						)
					}
				}
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name       string
		rows, cols int
		first      []float64
		second     []float64 // the dimensions of the second will be inverted
		expected   []float64
	}{
		{
			"1x1", 1, 1, []float64{15}, []float64{10}, []float64{150},
		},
		{
			"1x3 x 3x1 Vector", 1, 3,
			[]float64{1, 2, 3},
			[]float64{1, 2, 3},
			[]float64{14},
		},
		{
			"3x3 Matrices", 3, 3,
			[]float64{1, 3, 9, 2, 4, 6, 7, 14, 21},
			[]float64{0, 1, 0, 1, 0, 0, 1, 0, 0},
			[]float64{12, 1, 0, 10, 2, 0, 35, 7, 0},
		},
		{
			"3x3 Matrix Right Identity", 3, 3,
			[]float64{1, 3, 9, 2, 4, 6, 7, 14, 21},
			[]float64{1, 0, 0, 0, 1, 0, 0, 0, 1},
			[]float64{1, 3, 9, 2, 4, 6, 7, 14, 21},
		},
		{
			"3x3 Matrix Left Identity", 3, 3,
			[]float64{1, 0, 0, 0, 1, 0, 0, 0, 1},
			[]float64{1, 3, 9, 2, 4, 6, 7, 14, 21},
			[]float64{1, 3, 9, 2, 4, 6, 7, 14, 21},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := matrix.NewFromSlice(test.first, test.rows, test.cols)
			multiplier := matrix.NewFromSlice(test.second, test.cols, test.rows)
			result := matrix.Multiply(m, multiplier)
			index := 0
			// the dimensions for m x n inputs will always be n x n
			for row := 0; row < test.rows; row++ {
				for col := 0; col < test.rows; col++ {
					expected := test.expected[index]
					expectedResult := result.Get(row, col)
					if expected != expectedResult {
						t.Fatalf(
							"At row %d, col %d the matrix was expected to return %f "+
								"as prescribed in the test data (%v multiplied by %v). "+
								"Instead it returned %f",
							row,
							col,
							expected,
							test.first,
							test.second,
							result.Get(row, col),
						)
					}
					index++
				}
			}
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name       string
		rows, cols int
		data       []float64
		function   func(float64) float64
	}{
		{
			"1x1", 1, 1, []float64{15}, func(x float64) float64 { return x * 2 },
		},
		{
			"3x1 Vector", 3, 1, []float64{15, 64, 32}, func(x float64) float64 { return x * 3.15 },
		},
		{
			"3x3 Matrix", 3, 3, []float64{1, 3, 9, 2, 4, 6, 7, 14, 21}, func(x float64) float64 { return x * 1.28 },
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := matrix.NewFromSlice(test.data, test.rows, test.cols)
			mapped := matrix.Map(m, test.function)
			index := 0
			for row := 0; row < test.rows; row++ {
				for col := 0; col < test.cols; col++ {
					expectedResult := test.function(test.data[index])
					result := mapped.Get(row, col)
					if result != expectedResult {
						t.Fatalf(
							"At row %d, col %d the matrix was expected to return %f "+
								"as prescribed in the test data (%v)."+
								"Instead it returned %f",
							row,
							col,
							expectedResult,
							test.data,
							result,
						)
					}
					index++
				}
			}
		})
	}
}
