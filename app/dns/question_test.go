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

func TestQuestionsFromBytes(t *testing.T) {
	// only questions (headers already pruned)
	{
		expected := []Question{
			{Name: NameFromString("google.com"), Type: 1, Class: 1},
			{Name: NameFromString("booble.pl"), Type: 1, Class: 1},
		}
		expectedConsumed := uint(31)
		result, consumedBytes := QuestionsFromBytes([]byte("\x06google\x03com\x00\x00\x01\x00\x01\x06booble\x02pl\x00\x00\x01\x00\x01"), 2)

		if len(result) != len(expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}

		if expectedConsumed != consumedBytes {
			t.Errorf("expected %v, got %v", expectedConsumed, consumedBytes)
		}
	}

	// with other following RRs
	{
		expected := []Question{
			{Name: NameFromString("google.com"), Type: 1, Class: 1},
			{Name: NameFromString("booble.pl"), Type: 1, Class: 1},
		}
		expectedConsumed := uint(31)
		result, consumedBytes := QuestionsFromBytes([]byte("\x06google\x03com\x00\x00\x01\x00\x01\x06booble\x02pl\x00\x00\x01\x00\x01SOME_OTHER_BYTES_FOR_OTHER_RRS"), 2)

		if len(result) != len(expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}

		if expectedConsumed != consumedBytes {
			t.Errorf("expected %v, got %v", expectedConsumed, consumedBytes)
		}
	}
}
