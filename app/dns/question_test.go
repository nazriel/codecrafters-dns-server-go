package dns

import (
	"bytes"
	"testing"
)

func TestQuestionLabel(t *testing.T) {
	expected := []byte("\x06google\x03com\x00")
	result := QuestionNameFromString("google.com").AsLabel()

	if !bytes.Equal(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestQuestionBytes(t *testing.T) {
	expected := []byte("\x06google\x03com\x00\x00\x01\x00\x01")
	result := Question{Name: QuestionNameFromString("google.com"), Type: 1, Class: 1}.Bytes()

	if !bytes.Equal(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
