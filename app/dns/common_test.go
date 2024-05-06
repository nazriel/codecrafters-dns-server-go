package dns

import (
	"bytes"
	"testing"
)

func TestNameFromBytes(t *testing.T) {
	// ignore things after \0
	{
		result := NameFromBytes([]byte("\x06google\x03com\x00\x06booble\x02pl"))
		expected := Name{Parts: []NamePart{{"google", 6}, {"com", 3}}}
		if len(result.Parts) != len(expected.Parts) {
			t.Errorf("expected %v, got %v", expected, result)
		}

		if result.Parts[0].Str != expected.Parts[0].Str {
			t.Errorf("expected %v, got %v", expected, result)
		}
	}

	// weird domain google.com.booble.pl
	{
		result := NameFromBytes([]byte("\x06google\x03com\x06booble\x02pl\x00"))
		expected := Name{Parts: []NamePart{{"google", 6}, {"com", 3}, {"booble", 6}, {"pl", 2}}}
		if len(result.Parts) != len(expected.Parts) {
			t.Errorf("expected %v, got %v", expected, result)
		}

		if result.Parts[0].Str != expected.Parts[0].Str {
			t.Errorf("expected %v, got %v", expected, result)
		}
	}

	// unterminated
	{
		result := NameFromBytes([]byte("\x06google\x03com"))
		if len(result.Parts) != 0 {
			t.Errorf("expected empty, got %v", result)
		}
	}
}

func TestNameLabel(t *testing.T) {
	expected := []byte("\x06google\x03com\x00")
	result := NameFromString("google.com").AsLabel()

	if !bytes.Equal(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestIpFromString(t *testing.T) {
	{
		expected := uint32(134744072)
		result := IpFromString("8.8.8.8")

		if expected != result {
			t.Errorf("expected %v, got %v", expected, result)
		}
	}

	{
		result := IpFromString("256.255.255.255")

		if result != 0 {
			t.Errorf("expected %v, got %v", nil, result)
		}
	}
}
