package neural_test

import (
	"math"
	"testing"

	"example.com/m/internal/matrix"
	"example.com/m/internal/neural"
)

func sigmoid(x float64) float64 {
	return 1 / (1 + (math.Exp(-x)))
}

func TestPredict(t *testing.T) {
	inputSize := 10
	outputSize := 5
	ones := make([]float64, inputSize)
	for i := range ones {
		ones[i] = 1
	}
	input := matrix.NewFromSlice(ones, inputSize, 1)
	n := neural.New([]int{inputSize, outputSize}, sigmoid)
	output := n.Predict(input)
	t.Logf(
		"%s: network.Predict() successfully makes a prediction without erroring:\n%s",
		t.Name(),
		output,
	)
}
