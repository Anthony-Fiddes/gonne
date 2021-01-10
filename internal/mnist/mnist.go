// Package mnist provides functions to easily import mnist image set data
package mnist

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Error represents an error in reading mnist data
type Error string

func (e Error) Error() string {
	return string(e)
}

// ErrInvalidMagicNumber specifies that the data being read did not have the
// correct magic number.
const ErrInvalidMagicNumber Error = "mnist: invalid magic number"

const (
	unxepectedReadErr = "mnist: unexpected error while reading: %w"
	initErr           = "mnist: unexpected error while setting magic numbers: %w"
)

// Set is an MNIST set of images and labels
type Set struct {
	Labels []byte
	Images []image
}

type image struct {
	Rows   int32
	Cols   int32
	Pixels []byte
}

const (
	labelMagicNumber int32 = 0x801
	imageMagicNumber int32 = 0x803
)

var (
	byteOrder = binary.BigEndian
)

func readLabels(r io.Reader) ([]byte, error) {
	header := struct {
		Magic int32
		Size  int32
	}{}
	err := binary.Read(r, byteOrder, &header)
	if err != nil {
		return nil, fmt.Errorf(unxepectedReadErr, err)
	}
	if header.Magic != labelMagicNumber {
		return nil, ErrInvalidMagicNumber
	}

	labels := make([]byte, header.Size)
	err = binary.Read(r, byteOrder, labels)
	if err != nil {
		return nil, fmt.Errorf(unxepectedReadErr, err)
	}
	return labels, nil
}
