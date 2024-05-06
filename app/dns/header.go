package dns

import (
	"encoding/binary"
)

type HeaderFlags struct {
	QueryResponseIndiciator uint8 // 1bit
	Opcode                  uint8 // 4bit
	AuthoritativeAnswer     uint8 // 1 bit
	Truncation              uint8 // 1 bit
	RecursionDesired        uint8 // 1 bit
	RecursionAvailable      uint8 // 1 bit
	Reserved                uint8 // 3 bit
	ResponseCode            uint8 // 4 bit
}

func (hf HeaderFlags) Packed() uint16 {
	flags := uint16(hf.QueryResponseIndiciator) << 15
	flags += uint16(hf.Opcode) << 11
	flags += uint16(hf.AuthoritativeAnswer) << 10
	flags += uint16(hf.Truncation) << 9
	flags += uint16(hf.RecursionDesired) << 8
	flags += uint16(hf.RecursionAvailable) << 7
	flags += uint16(hf.Reserved) << 4
	flags += uint16(hf.ResponseCode)
	return flags
}

type Header struct {
	PacketID              uint16
	HeaderFlags           HeaderFlags
	QuestionCount         uint16
	AnswerRecordCount     uint16
	AuthorityRecordCount  uint16
	AdditionalRecordCount uint16
}

func (header Header) Bytes() []byte {
	buf := make([]byte, 12)
	binary.BigEndian.PutUint16(buf[:2], header.PacketID)
	binary.BigEndian.PutUint16(buf[2:4], header.HeaderFlags.Packed())
	binary.BigEndian.PutUint16(buf[4:6], header.QuestionCount)
	binary.BigEndian.PutUint16(buf[6:8], header.AnswerRecordCount)
	binary.BigEndian.PutUint16(buf[8:10], header.AuthorityRecordCount)
	binary.BigEndian.PutUint16(buf[10:], header.AdditionalRecordCount)

	return buf
}
