package dns

import (
	"encoding/binary"
)

type Question struct {
	Name  Name
	Type  uint16
	Class uint16
}

func (q Question) Bytes() []byte {
	buf := q.Name.AsLabel()
	buf = binary.BigEndian.AppendUint16(buf, q.Type)
	buf = binary.BigEndian.AppendUint16(buf, q.Class)
	return buf
}

func QuestionsFromBytes(payload []byte, expectedLen uint16) ([]Question, uint) {
	questions := []Question{}
	consumedTotal := uint(0)

	for foundQ := uint16(0); foundQ < expectedLen; foundQ++ {
		question := Question{}

		name := NameFromBytes(payload)
		if len(name.Parts) == 0 {
			return nil, consumedTotal
		}
		consumed := uint(0)
		for _, part := range name.Parts {
			consumed += 1 // len
			consumed += uint(len(part.Str))
		}
		question.Name = name
		consumed += 1 // null terminator

		question.Type = binary.BigEndian.Uint16(payload[consumed : consumed+2])
		consumed += 2

		question.Class = binary.BigEndian.Uint16(payload[consumed : consumed+2])
		consumed += 2

		consumedTotal += consumed
		payload = payload[consumed:]

		questions = append(questions, question)
	}

	return questions, uint(consumedTotal)
}
