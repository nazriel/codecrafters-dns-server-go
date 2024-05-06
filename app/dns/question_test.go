package dns

import (
	"bytes"
	"testing"
)

func TestQuestionBytes(t *testing.T) {
	expected := []byte("\x06google\x03com\x00\x00\x01\x00\x01")
	result := Question{Name: NameFromString("google.com"), Type: 1, Class: 1}.Bytes()

	if !bytes.Equal(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
