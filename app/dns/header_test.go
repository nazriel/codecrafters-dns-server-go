package dns

import (
	"bytes"
	"testing"
)

func TestHeaderFlagsPacked(t *testing.T) {
	{
		result := HeaderFlags{1, 1, 1, 1, 1, 1, 1, 1}.Packed()
		expected := uint16(0b1000111110010001)

		if result != expected {
			t.Errorf("got %d, expected %d", result, expected)
		}
	}

	{
		result := HeaderFlags{1, 1, 1, 1, 0, 0, 0, 0}.Packed()
		expected := uint16(0b1000111000000000)

		if result != expected {
			t.Errorf("got %d, expected %d", result, expected)
		}
	}

	{
		result := HeaderFlags{QueryResponseIndiciator: 1}.Packed()
		expected := uint16(0b1000000000000000)

		if result != expected {
			t.Errorf("got %d, expected %d", result, expected)
		}
	}
}

func TestHeaderBytes(t *testing.T) {
	{
		result := Header{
			PacketID:    1234,
			HeaderFlags: HeaderFlags{QueryResponseIndiciator: 1},
		}.Bytes()

		expected := []byte{0x04, 0xd2, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0}

		if bytes.Equal(result, expected) == false {
			t.Errorf("got %d, expected %d", result, expected)
		}

	}
}
