// Package neural supplies neural network operations
//
// Created with reference to Sebastian Lague's amazing youtube series on neural
// networks. 3Blue1Brown is also a great channel to reference.
package neural

import "example.com/m/internal/matrix"

// Network represents a neural network
type Network struct {
	layerSizes []int
	weights    []*matrix.Matrix
	biases     []*matrix.Matrix
	Activation func(float64) float64
}

// New returns a new neural network
func New(layerSizes []int, activation func(float64) float64) *Network {
	// There has to be at least two layers for input and output
	if len(layerSizes) < 2 {
		panic("neural: there must be at least 2 layers (one for input and one for output)")
	}

	net := Network{layerSizes: layerSizes}

	netLayerSizes := make([]int, len(layerSizes))
	copy(netLayerSizes, layerSizes)
	net.layerSizes = netLayerSizes

	net.weights = make([]*matrix.Matrix, 0, len(net.layerSizes)-1)
	net.biases = make([]*matrix.Matrix, 0, len(net.layerSizes)-1)
	for i := 1; i < len(net.layerSizes); i++ {
		rows := net.layerSizes[i]
		cols := net.layerSizes[i-1]
		net.weights = append(net.weights, matrix.NewRandomNormal(rows, cols))
		net.biases = append(net.biases, matrix.New(rows, 1))
	}

	net.Activation = activation

	return &net
}

// Predict takes an input matrix and produces a matrix describing the
// probabilities for each possible output.
func (n *Network) Predict(input *matrix.Matrix) *matrix.Matrix {
	result := input
	for i := 0; i < len(n.weights); i++ {
		result = matrix.Multiply(n.weights[i], result)
		result = matrix.Add(result, n.biases[i])
		result = matrix.Map(result, n.Activation)
	}
	return result
}
