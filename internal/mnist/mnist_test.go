package mnist

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"testing"
)

func digits() []byte {
	return []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
}

func createTestLabelData(labels []byte) (data *bytes.Buffer) {
	size := uint32(len(labels))
	s := make([]byte, 0, size+8)
	testData := bytes.NewBuffer(s)

	binary.Write(testData, byteOrder, labelMagicNumber)
	binary.Write(testData, byteOrder, size)
	binary.Write(testData, byteOrder, labels)
	return testData
}

func createTestImageData(images []Image) (data *bytes.Buffer) {
	size := uint32(len(images))
	testData := &bytes.Buffer{}

	binary.Write(testData, byteOrder, imageMagicNumber)
	binary.Write(testData, byteOrder, size)
	for _, image := range images {
		binary.Write(testData, byteOrder, image.Rows)
		binary.Write(testData, byteOrder, image.Cols)
		binary.Write(testData, byteOrder, image.Pixels)
	}
	return testData
}

// TODO: Add more test cases
// TODO: Make error messages more robust

func TestReadLabels(t *testing.T) {
	expected := digits()
	testLabelSet := createTestLabelData(expected)

	labels, err := ReadLabels(testLabelSet)
	if err != nil {
		t.Fatal(err)
	}

	for i := range expected {
		if expected[i] != labels[i] {
			t.Fatalf("Expected %v but got %v", expected, labels)
		}
	}
}

func TestReadImages(t *testing.T) {
	expected := []Image{
		{1, 1, []byte{0}},
		{2, 2, []byte{0, 1, 0, 1}},
	}
	testImageSet := createTestImageData(expected)

	images, err := ReadImages(testImageSet)
	if err != nil {
		t.Fatal(err)
	}

	for i := range expected {
		if !reflect.DeepEqual(expected[i], images[i]) {
			t.Fatalf("Expected %v but got %v", expected, images)
		}
	}
}
