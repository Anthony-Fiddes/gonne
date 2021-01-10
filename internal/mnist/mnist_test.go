package mnist

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func digits() []byte {
	return []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
}

func createTestLabelSet() *bytes.Buffer {
	labels := digits()
	size := uint32(len(labels))
	s := make([]byte, 0, 11)
	testData := bytes.NewBuffer(s)

	binary.Write(testData, byteOrder, labelMagicNumber)
	binary.Write(testData, byteOrder, size)
	binary.Write(testData, byteOrder, labels)
	return testData
}

func TestReadLabels(t *testing.T) {
	testLabelSet := createTestLabelSet()
	expected := digits()

	labels, err := readLabels(testLabelSet)
	if err != nil {
		t.Fatal(err)
	}

	for i := range expected {
		if expected[i] != labels[i] {
			t.Fatalf("Expected %v but got %v", expected, labels)
		}
	}
}
