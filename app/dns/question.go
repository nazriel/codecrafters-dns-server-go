package dns

import (
	"encoding/binary"
	"strings"
)

type QuestionNamePart struct {
	Name   string
	Length uint8
}

type QuestionName struct {
	Parts []QuestionNamePart
}

func (qn QuestionName) AsLabel() []byte {
	result := []byte{}
	for _, part := range qn.Parts {
		result = append(result, part.Length)
		result = append(result, []byte(part.Name)...)
	}
	result = append(result, 0)

	return result
}

func QuestionNameFromString(domain string) QuestionName {
	qn := QuestionName{}

	for _, part := range strings.Split(domain, ".") {
		qn.Parts = append(qn.Parts, QuestionNamePart{part, uint8(len(part))})
	}

	return qn
}

type Question struct {
	Name  QuestionName
	Type  uint16
	Class uint16
}

func (q Question) Bytes() []byte {
	buf := q.Name.AsLabel()
	buf = binary.BigEndian.AppendUint16(buf, q.Type)
	buf = binary.BigEndian.AppendUint16(buf, q.Class)
	return buf
}
