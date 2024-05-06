package dns

import (
	"bytes"
	"testing"
)

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
