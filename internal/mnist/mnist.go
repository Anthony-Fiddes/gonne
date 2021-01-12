// Package mnist provides functions to easily import mnist data
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
	Images []Image
}

// Image is an MNIST image
type Image struct {
	Rows   int32
	Cols   int32
	Pixels []byte
}

// Get returns the pixel at the given row and column
func (i *Image) Get(row, col int) byte {
	return i.Pixels[int(i.Cols)*row+col]
}

const (
	labelMagicNumber int32 = 0x801
	imageMagicNumber int32 = 0x803
)

var (
	byteOrder = binary.BigEndian
)

func readHeader(r io.Reader) (magic, size int32, err error) {
	header := struct {
		Magic int32
		Size  int32
	}{}
	err = binary.Read(r, byteOrder, &header)
	if err != nil {
		return 0, 0, err
	}
	return header.Magic, header.Size, nil
}

// ReadLabels reads the labels file of an MNIST data set
func ReadLabels(r io.Reader) ([]byte, error) {
	magic, size, err := readHeader(r)
	if err != nil {
		return nil, fmt.Errorf(unxepectedReadErr, err)
	}
	if magic != labelMagicNumber {
		return nil, ErrInvalidMagicNumber
	}

	labels := make([]byte, size)
	err = binary.Read(r, byteOrder, labels)
	if err != nil {
		return nil, fmt.Errorf(unxepectedReadErr, err)
	}
	return labels, nil
}

// ReadImages reads the images file of an MNIST data set
func ReadImages(r io.Reader) ([]Image, error) {
	magic, size, err := readHeader(r)
	if err != nil {
		return nil, fmt.Errorf(unxepectedReadErr, err)
	}
	if magic != imageMagicNumber {
		return nil, ErrInvalidMagicNumber
	}

	imageHeader := struct {
		Rows int32
		Cols int32
	}{}
	readImage := func() (Image, error) {
		err = binary.Read(r, byteOrder, &imageHeader)
		if err != nil {
			return Image{}, err
		}

		pixels := make([]byte, int(imageHeader.Rows*imageHeader.Cols))
		err = binary.Read(r, byteOrder, pixels)
		if err != nil {
			return Image{}, err
		}
		return Image{imageHeader.Rows, imageHeader.Cols, pixels}, nil
	}

	images := make([]Image, 0, size)
	for i := 0; i < int(size); i++ {
		image, err := readImage()
		if err != nil {
			return nil, fmt.Errorf(unxepectedReadErr, err)
		}
		images = append(images, image)
	}
	return images, nil
}
