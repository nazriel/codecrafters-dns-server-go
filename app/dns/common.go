package dns

import (
	"encoding/binary"
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
