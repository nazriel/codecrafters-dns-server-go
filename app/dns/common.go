package dns

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

type NamePart struct {
	Str    string
	Length uint8
}

type Name struct {
	Parts []NamePart
}

func (qn Name) AsLabel() []byte {
	result := []byte{}
	for _, part := range qn.Parts {
		result = append(result, part.Length)
		result = append(result, []byte(part.Str)...)
	}
	result = append(result, 0)

	return result
}

func NameFromBytes(payload []byte) Name {
	// len, str, len, str, 0
	result := Name{}
	for {
		if len(payload) == 0 {
			fmt.Println("unterminated labels")
			return Name{}
		}
		length := payload[0]
		if length == 0 {
			break
		}
		if length < 1 || length > 255 {
			fmt.Println("wrong label length")
			return Name{}
		}
		str := payload[1 : length+1]

		result.Parts = append(result.Parts, NamePart{string(str), length})
		payload = payload[length+1:]
	}

	return result
}

func NameFromString(domain string) Name {
	qn := Name{}

	for _, part := range strings.Split(domain, ".") {
		qn.Parts = append(qn.Parts, NamePart{part, uint8(len(part))})
	}

	return qn
}

func IpFromString(ip string) uint32 {
	octets := strings.Split(ip, ".")
	if len(octets) != 4 {
		return 0
	}

	var result []byte
	for _, octet := range octets {
		if number, err := strconv.ParseUint(octet, 10, 8); err == nil {
			result = append(result, byte(number))
		} else {
			return 0
		}
	}

	return binary.BigEndian.Uint32(result[:])
}
